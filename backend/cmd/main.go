// Package main это входная точка проекта
package main

import (
	"backend/cmd/app"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, _ []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	app := app.NewApp(ctx, w)
	err := app.Run()
	if err != nil {
		return err
	}

	return nil
}
