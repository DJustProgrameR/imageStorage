// Package repository это доступ к изображению на файловой системе
package repository

import (
	"backend/internal/model"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	petPathStorageType string = ".txt"
)

// Config of repo
type Config interface {
	GetPetStoragePath() string
}

// NewPetRepository inits repo
func NewPetRepository(config Config) *PetRepository {
	if config == nil {
		log.Fatalf("PetRepository config is null")
	}
	path := config.GetPetStoragePath()
	if path == "" || !strings.HasSuffix(path, petPathStorageType) {
		log.Fatalf("PetRepository invalid config path %s", path)
	}
	return &PetRepository{
		petPath: path,
	}
}

// PetRepository interface
type PetRepository struct {
	petPath           string
	createStorageOnce sync.Once
}

// GetPet gets pet from storage
func (r *PetRepository) GetPet(ctx context.Context) (pet *model.Pet, err error) {
	r.createStorageOnce.Do(func() {
		err = r.initFile(r.petPath)
	})
	if err != nil {
		return
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	pet = &model.Pet{}

	data, err := os.ReadFile(r.petPath)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, pet)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return pet, nil
}

// PutPet puts pet to storage
func (r *PetRepository) PutPet(ctx context.Context, pet model.Pet) (err error) {
	r.createStorageOnce.Do(func() {
		err = r.initFile(r.petPath)
	})
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(pet, "", "  ")
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	err = os.WriteFile(r.petPath, data, 0600)
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	return nil
}

func (r *PetRepository) initFile(filePath string) (err error) {
	if _, err := os.Stat(filePath); err != nil && !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()

	return nil
}
