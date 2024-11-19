package main

import (
	"context"
	"log/slog"

	"github.com/marcelofabianov/cashly/config"
	app "github.com/marcelofabianov/cashly/internal"
	"github.com/marcelofabianov/cashly/pkg/database"
)

func main() {
	run()

	slog.Info("API is running")
}

func run() error {
	// Config
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
	}

	// Database
	ctx := context.Background()
	db, err := database.Connect(ctx, cfg.Db)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("error closing database connection", "error", err)
		}
	}()

	// App
	app := app.NewApp(cfg, db.Conn())
	app.Build()

	return nil
}
