// Package app provides runnable app
package app

import (
	"backend/cmd/config"
	"backend/internal/repository"
	"backend/internal/server"
	"backend/internal/usecase"
	"context"
	"io"
	"net/http"
)

// NewApp provides main backend handler
func NewApp(ctx context.Context, w io.Writer) *App {
	app := &App{}

	appConfig := config.NewAppConfig()
	app.port = ":" + appConfig.GetAppPort()

	petRepo := repository.NewPetRepository(appConfig)

	putPetUseCase := usecase.NewPutPet(w, petRepo)
	getPetUseCase := usecase.NewGetPet(petRepo)

	restServer := server.NewRestServer(ctx, w, getPetUseCase, putPetUseCase)
	serverWrapper := server.Handler(restServer)

	app.handler = serverWrapper

	return app
}

type App struct {
	port    string
	handler http.Handler
}

func (a *App) Run() error {
	err := http.ListenAndServe(a.port, a.handler)
	if err != nil {
		return err
	}

	return nil
}
