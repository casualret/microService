package service

import (
	"log/slog"
	"urlshortener/internal/storage"
)

type App struct {
	BannerOperations
	ParamOperations
}

func NewApp(logger *slog.Logger, storage *storage.Postgres) (*App, error) {
	slog.Info("Taalk")
	return &App{BannerOperations: NewBannerManager(storage),
		ParamOperations: NewParamOpManager(storage),
	}, nil
}
