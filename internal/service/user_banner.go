package service

import (
	"fmt"
	"urlshortener/internal/models"
	"urlshortener/internal/storage"
)

type UBannerOperations interface {
	GetUserBanner(req models.GetUserBannerReq) (*models.BannerWithDetails, error)
}

type UBannerManager struct {
	storage *storage.Postgres
}

func NewUBannerManager(storage *storage.Postgres) *UBannerManager {
	return &UBannerManager{storage: storage}
}

func (ubm *UBannerManager) GetUserBanner(req models.GetUserBannerReq) (*models.BannerWithDetails, error) {
	const op = "service.GetUserBanner"

	banner, err := ubm.storage.GetUserBanner(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return banner, nil
}
