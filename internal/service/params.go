package service

import (
	"fmt"
	"urlshortener/internal/models"
	"urlshortener/internal/storage"
)

type ParamOperations interface {
	CreateTag(tag models.Tag) error
	CreateFeature(feature models.Feature) error
}

type ParamOpManager struct {
	storage *storage.Postgres
}

func NewParamOpManager(storage *storage.Postgres) *ParamOpManager {
	return &ParamOpManager{storage: storage}
}

func (pm *ParamOpManager) CreateTag(tag models.Tag) error {
	const op = "service.CreateTag"

	err := pm.storage.CreateTag(tag)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pm *ParamOpManager) CreateFeature(feature models.Feature) error {
	const op = "service.CreateFeature"

	err := pm.storage.CreateFeature(feature)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
