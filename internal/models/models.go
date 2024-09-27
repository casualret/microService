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
	TagIds    []int64                `json:"tag_ids,omitempty"`
	FeatureID int64                  `json:"feature_id,omitempty"`
	NewBanner map[string]interface{} `json:"content" binding:"required"`
	IsActive  bool                   `json:"is_active" binding:"required"`
}

type ChangeBannerReq struct {
	TagIds    []int64                `json:"tag_ids,omitempty"`
	FeatureID int64                  `json:"feature_id,omitempty"`
	NewBanner map[string]interface{} `json:"content,omitempty"`
	IsActive  *bool                  `json:"is_active,omitempty"`
}

type GetBannersReq struct {
	FeatureID *int
	TagID     *int
	Limit     *int
	Offset    *int
}

type GetUserBannerReq struct {
	TagID           string
	FeatureID       string
	UseLastRevision bool
}
