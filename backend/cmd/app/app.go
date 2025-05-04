// Package app provides runnable app
package app

import (
	"backend/cmd/config"
	"backend/internal/repository"
	"backend/internal/server"
	"backend/internal/usecase"
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

// NewApp provides main backend handler
func NewApp(ctx context.Context, w io.Writer) *App {
	logger := log.New(w, "ERROR: ", log.LstdFlags)
	app := &App{}

	appConfig := config.NewAppConfig()
	app.port = ":" + appConfig.GetAppPort()

	petRepo := repository.NewPetRepository(appConfig)

	putPetUseCase := usecase.NewPutPet(logger, petRepo)
	getPetUseCase := usecase.NewGetPet(logger, petRepo)

	restServer := server.NewRestServer(ctx, logger, getPetUseCase, putPetUseCase)
	serverWrapper := server.Handler(restServer)

	app.handler = serverWrapper

	return app
}

// App -
type App struct {
	port    string
	handler http.Handler
}

// Run runs app on port
func (a *App) Run() error {

	restServer := &http.Server{
		Addr:         a.port,
		Handler:      a.handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := restServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
