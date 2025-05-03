// Package usecase это бизнес логика
package usecase

import (
	"backend/internal/model"
	"context"
)

// PetRepository interface
type PetRepository interface {
	GetPet(ctx context.Context) (*model.Pet, error)
	PutPet(ctx context.Context, pet model.Pet) error
}

func NewGetPet(repo PetRepository) *GetPet {
	return &GetPet{repo: repo}
}

type GetPet struct {
	repo PetRepository
}

func (uc *GetPet) Execute(ctx context.Context) (*model.Pet, error) {
	select {
	case <-ctx.Done():
		return nil, model.ErrInternal
	default:
		pet, err := uc.repo.GetPet(ctx)
		if err != nil {
			return nil, model.ErrInternal
		}
		if pet == nil {
			return nil, model.ErrNoData
		}
		return pet, nil
	}
}
