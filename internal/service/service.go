package service

import (
	"context"
	"log/slog"
	"microService/internal/redis_cash"
	"microService/internal/storage"
)

type App struct {
	Authentication
	BannerOperations
	ParamOperations
	UBannerOperations
}

func NewApp(logger *slog.Logger, storage *storage.Postgres, cash *redis_cash.RedisCash, ctx *context.Context) (*App, error) {
	slog.Info("Taalk")
	return &App{Authentication: NewAuthenticationService(storage),
		BannerOperations:  NewBannerManager(storage, cash, ctx),
		ParamOperations:   NewParamOpManager(storage),
		UBannerOperations: NewUBannerManager(storage, cash),
	}, nil
}
