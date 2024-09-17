package service

import (
	"context"
	"log/slog"
	"urlshortener/internal/models"
	"urlshortener/internal/storage"
)

type App struct {
	BannerOperations
	ParamOperations
}

type BannerOperations interface {
	//CreateBanner(ctx context.Context, req models.CreateBannerReq) error
	//ChangeBanner(ctx context.Context, bannerID int54, req models.ChangeBannerReq) error
	//GetBanners(ctx context.Context, req models.GetAllBannersParams) ([]*models.BannerWithDetails, error)
}

type BannerManager struct {
	storage *storage.Postgres
}

func NewBannerManager(storage *storage.Postgres) *BannerManager {
	return &BannerManager{storage: storage}
}

func (bm *BannerManager) CreateBanner(ctx context.Context, req models.CreateBannerReq) error {
	_, _ = ctx, req
	return nil
}

func NewApp(logger *slog.Logger, storage *storage.Postgres) (*App, error) {
	slog.Info("Taalk")
	return &App{BannerOperations: NewBannerManager(storage),
		ParamOperations: NewParamOpManager(storage),
	}, nil
}
