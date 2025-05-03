// Package server это конфигурация REST API сервера
package server

import (
	"backend/internal/model"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type GetUseCase interface {
	Execute(context.Context) (*model.Pet, error)
}

type PutUseCase interface {
	Execute(ctx context.Context, pet model.Pet) error
}

// NewRestServer -
func NewRestServer(
	ctx context.Context,
	w io.Writer,
	getUseCase GetUseCase,
	putUseCase PutUseCase,
) *RestServer {
	return &RestServer{
		ctx:        ctx,
		w:          w,
		getUseCase: getUseCase,
		putUseCase: putUseCase,
	}
}

// NewRestServer implementation
type RestServer struct {
	ctx        context.Context
	w          io.Writer
	getUseCase GetUseCase
	putUseCase PutUseCase
}

// GetPet -
func (s *RestServer) GetPet(w http.ResponseWriter, _ *http.Request) {
	select {
	case <-s.ctx.Done():
		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.Header().Set("Content-Type", "application/json")
		pet, err := s.getUseCase.Execute(s.ctx)

		if err != nil {
			if err == model.ErrInternal {
				w.WriteHeader(http.StatusInternalServerError)
				err = json.NewEncoder(w).Encode(err)
				return
			} else if err == model.ErrNoData {
				w.WriteHeader(http.StatusNoContent)
				return
			} else {
				w.WriteHeader(http.StatusBadRequest)
				err = json.NewEncoder(w).Encode(err)
				if err != nil {
					s.w.Write([]byte("GetPet" + err.Error()))
					return
				}
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(pet)
		if err != nil {
			s.w.Write([]byte("GetPet" + err.Error()))
			return
		}
	}
}

// UploadPet -
func (s *RestServer) UploadPet(w http.ResponseWriter, r *http.Request) {
	select {
	case <-s.ctx.Done():
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	default:
		w.Header().Set("Content-Type", "application/json")
		request := model.Pet{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.putUseCase.Execute(s.ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(err)
			if err != nil {
				s.w.Write([]byte("PutPet" + err.Error()))
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}

}
