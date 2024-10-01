package service

import (
	"context"
	"log/slog"
	"microService/internal/storage"
)

type App struct {
	Authentication
	BannerOperations
	ParamOperations
	UBannerOperations
	Ctx *context.Context
}

func NewApp(logger *slog.Logger, storage *storage.Postgres, ctx *context.Context) (*App, error) {
	slog.Info("Taalk")
	return &App{Authentication: NewAuthenticationService(storage),
		BannerOperations:  NewBannerManager(storage),
		ParamOperations:   NewParamOpManager(storage),
		UBannerOperations: NewUBannerManager(storage),
		Ctx:               ctx,
	}, nil
}
