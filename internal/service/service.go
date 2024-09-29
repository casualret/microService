package service

import (
	"log/slog"
	"microService/internal/storage"
)

type App struct {
	Authentication
	BannerOperations
	ParamOperations
	UBannerOperations
}

func NewApp(logger *slog.Logger, storage *storage.Postgres) (*App, error) {
	slog.Info("Taalk")
	return &App{Authentication: NewAuthenticationService(storage),
		BannerOperations:  NewBannerManager(storage),
		ParamOperations:   NewParamOpManager(storage),
		UBannerOperations: NewUBannerManager(storage),
	}, nil
}
