// Package usecase это бизнес логика
package usecase

import (
	"backend/internal/model"
	"context"
	"log"
)

// PetRepository interface
type PetRepository interface {
	GetPet(ctx context.Context) (*model.Pet, error)
	PutPet(ctx context.Context, pet model.Pet) error
}

// NewGetPet returns usecase
func NewGetPet(
	logger *log.Logger,
	repo PetRepository,
) *GetPet {
	if logger == nil {
		log.Fatalf("GetPet logger w is null")
	}
	if repo == nil {
		log.Fatalf("GetPet usecase repo is null")
	}
	return &GetPet{
		logger: logger,
		repo:   repo,
	}
}

// GetPet usecase
type GetPet struct {
	logger *log.Logger
	repo   PetRepository
}

// Execute usecase
func (uc *GetPet) Execute(ctx context.Context) (*model.Pet, error) {
	select {
	case <-ctx.Done():
		uc.logger.Printf("put pet context done: %s", ctx.Err().Error())
		return nil, model.ErrInternal
	default:
		pet, err := uc.repo.GetPet(ctx)
		if err != nil {
			uc.logger.Printf("get pet execute repo err: %s", err.Error())
			return nil, model.ErrInternal
		}
		if pet == nil {
			return nil, model.ErrNoData
		}
		return pet, nil
	}
}
