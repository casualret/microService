package service

import (
	"fmt"
	"urlshortener/internal/models"
	"urlshortener/internal/storage"
)

type BannerOperations interface {
	CreateBanner(req models.CreateBannerReq) error
	//ChangeBanner(ctx context.Context, bannerID int54, req models.ChangeBannerReq) error
	//GetBanners(ctx context.Context, req models.GetAllBannersParams) ([]*models.BannerWithDetails, error)
}

type BannerManager struct {
	storage *storage.Postgres
}

func NewBannerManager(storage *storage.Postgres) *BannerManager {
	return &BannerManager{storage: storage}
}

func (bm *BannerManager) CreateBanner(req models.CreateBannerReq) error {
	const op = "service.CreateBanner"

	err := bm.storage.CreateBanner(req)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
