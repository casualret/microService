package service

import (
	"fmt"
	"microService/internal/models"
	"microService/internal/redis_cash"
	"microService/internal/storage"
)

type UBannerOperations interface {
	GetUserBanner(req models.GetUserBannerReq) (*models.BannerWithDetails, error)
}

type UBannerManager struct {
	storage *storage.Postgres
	cash    *redis_cash.RedisCash
}

func NewUBannerManager(storage *storage.Postgres, cash *redis_cash.RedisCash) *UBannerManager {
	return &UBannerManager{storage: storage, cash: cash}
}

func (ubm *UBannerManager) GetUserBanner(req models.GetUserBannerReq) (*models.BannerWithDetails, error) {
	const op = "service.GetUserBanner"

	banner, err := ubm.storage.GetUserBanner(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return banner, nil
}
