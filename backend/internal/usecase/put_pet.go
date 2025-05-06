package usecase

import (
	"backend/internal/model"
	"context"
	"log"
	"strconv"
)

const (
	maxDescrLen = 300
	maxASCIILen = 10000
)

// NewPutPet returns usecase
func NewPutPet(logger *log.Logger, repo PetRepository) *PutPet {
	if logger == nil {
		log.Fatalf("PutPet usecase w is null")
	}
	if repo == nil {
		log.Fatalf("PutPet usecase repo is null")
	}

	return &PutPet{
		logger: logger,
		repo:   repo,
	}
}

// PutPet usecase
type PutPet struct {
	logger *log.Logger
	repo   PetRepository
}

// Execute usecase
func (uc *PutPet) Execute(ctx context.Context, pet model.Pet) error {
	select {
	case <-ctx.Done():
		uc.logger.Printf("put pet context done: %s", ctx.Err().Error())

		return model.ErrInternal
	default:
		err := uc.validate(pet)
		if err != nil {
			return err
		}
		err = uc.repo.PutPet(ctx, pet)
		if err != nil {
			uc.logger.Printf("put pet execute repo err: %s", err.Error())
			return model.ErrInternal
		}
		return nil
	}
}

func (uc *PutPet) validate(pet model.Pet) error {
	if len(pet.ASCII) > maxASCIILen {
		errMessage := model.ErrInvalidASCII
		errMessage.Message = errMessage.Message + " limit is " + strconv.Itoa(maxASCIILen)
		return errMessage
	}
	if len(pet.Description) > maxDescrLen {
		uc.logger.Printf("descr too long: %s", pet.Description)
		errMessage := model.ErrInvalidDescr
		errMessage.Message = errMessage.Message + " limit is " + strconv.Itoa(maxDescrLen)
		return errMessage
	}
	return nil
}
