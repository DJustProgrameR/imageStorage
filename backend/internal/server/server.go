package server

import (
	"backend/internal/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetUseCase -
type GetUseCase interface {
	Execute(context.Context) (*model.Pet, error)
}

// PutUseCase -
type PutUseCase interface {
	Execute(ctx context.Context, pet model.Pet) error
}

// NewRestServer -
func NewRestServer(
	ctx context.Context,
	logger *log.Logger,
	getUseCase GetUseCase,
	putUseCase PutUseCase,
) *RestServer {
	if getUseCase == nil {
		log.Fatalf("NewRestServer getUseCase is null")
	}
	if putUseCase == nil {
		log.Fatalf("NewRestServer putUseCase is null")
	}
	if ctx == nil {
		log.Fatalf("NewRestServer ctx is null")
	}
	if logger == nil {
		log.Fatalf("NewRestServer logger is null")
	}
	return &RestServer{
		ctx:        ctx,
		logger:     logger,
		getUseCase: getUseCase,
		putUseCase: putUseCase,
	}
}

// RestServer implementation
type RestServer struct {
	ctx        context.Context
	logger     *log.Logger
	getUseCase GetUseCase
	putUseCase PutUseCase
}

// GetPet -
func (s *RestServer) GetPet(w http.ResponseWriter, _ *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.Printf("Internal error: %v", err)
		}
	}()
	select {
	case <-s.ctx.Done():
		s.logger.Printf("RestServer GetPet context done: %s", s.ctx.Err().Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.Header().Set("Content-Type", "application/json")
		pet, err := s.getUseCase.Execute(s.ctx)

		if err != nil {
			if model.ErrInternal.Is(err) {
				w.WriteHeader(http.StatusInternalServerError)
				err = json.NewEncoder(w).Encode(err)
				s.logger.Printf("RestServer GetPet err encoding err: %s", err.Error())
				return
			} else if model.ErrNoData.Is(err) {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err)

			if err != nil {
				s.logger.Printf("RestServer UploadPet encode err err: %s", err.Error())
			}

			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(pet)
		if err != nil {
			s.logger.Printf("RestServer UploadPet encode pet err: %s", err.Error())
			return
		}
	}
}

// UploadPet -
func (s *RestServer) UploadPet(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.Printf("Internal error: %v", err)
		}
	}()
	select {
	case <-s.ctx.Done():
		s.logger.Printf("RestServer UploadPet context done: %s", s.ctx.Err().Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	default:
		w.Header().Set("Content-Type", "application/json")
		request := model.Pet{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &request)
		if err != nil {
			s.logger.Printf("RestServer UploadPet unmarsh err: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.putUseCase.Execute(s.ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(err)
			if err != nil {
				s.logger.Printf("RestServer UploadPet encode err err: %s", err.Error())
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}

}
