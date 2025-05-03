// Package usecase это бизнес логика
package usecase

import (
	"backend/internal/model"
	"context"
	"io"
)

const (
	maxDescrLen = 300
	maxAsciiLen = 10000
)

func NewPutPet(w io.Writer, repo PetRepository) *PutPet {
	return &PutPet{
		w:    w,
		repo: repo,
	}
}

type PutPet struct {
	w    io.Writer
	repo PetRepository
}

func (uc *PutPet) Execute(ctx context.Context, pet model.Pet) error {
	select {
	case <-ctx.Done():
		return model.ErrInternal
	default:
		//uc.w.Write([]byte("PutPetUseCase Execute: " + pet.Ascii))
		err := uc.validate(pet)
		if err != nil {
			return err
		}
		err = uc.repo.PutPet(ctx, pet)
		if err != nil {
			return model.ErrInternal
		}
		return nil
	}
}

func (uc *PutPet) validate(pet model.Pet) error {
	if len(pet.Ascii) > maxAsciiLen {
		return model.ErrInvalidAscii
	}
	if len(pet.Description) > maxDescrLen {
		return model.ErrInvalidAscii
	}
	return nil
}
