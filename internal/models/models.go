package models

import (
	"time"
)

type UserBanner struct {
	TagID           string
	FeatureID       string
	UseLastRevision bool
}

type BannerWithDetails struct {
	BannerID  int
	FeatureID *int
	TagIDs    []int64
	Content   map[string]interface{}
	Status    *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Tag struct {
	Name string `json:"name"`
}

type Feature struct {
	Name string `json:"name"`
}

type CreateBannerReq struct {
	TagIds    []string               `json:"tag_ids"`
	FeatureID string                 `json:"feature_id"`
	NewBanner map[string]interface{} `json:"new_banner"`
	IsActive  string                 `json:"is_active"`
}
