package models

import "time"

type UserBanner struct {
	TagID           string
	FeatureID       string
	UseLastRevision bool
}

type CreateBannerReq struct {
	TagIDs    []string
	FeatureID string
	NewBanner *map[string]interface{}
	IsActive  string
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
