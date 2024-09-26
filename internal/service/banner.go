package service

import (
	"fmt"
	"urlshortener/internal/models"
	"urlshortener/internal/storage"
)

type BannerOperations interface {
	CreateBanner(req models.CreateBannerReq) error
	GetBanners(req models.GetBannersReq) ([]*models.BannerWithDetails, error)
	DeleteBanner(bannerID int64) error
	//ChangeBanner(bannerID int64, req models.ChangeBannerReq) error
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

func (bm *BannerManager) GetBanners(req models.GetBannersReq) ([]*models.BannerWithDetails, error) {
	const op = "service.GetBanners"

	banners, err := bm.storage.GetBannersParams(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return banners, nil
}

func (bm *BannerManager) DeleteBanner(bannerID int64) error {
	const op = "service.DeleteBanner"

	err := bm.storage.DeleteBanner(bannerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
