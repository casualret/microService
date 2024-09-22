package storage

import (
	"encoding/json"
	"fmt"
	"urlshortener/internal/models"
)

func (p *Postgres) CreateBanner(banner models.CreateBannerReq) error {
	const op = "postgres.CreateBanner"

	BannerContent, err := json.Marshal(banner.NewBanner)

	tx, err := p.database.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	query := "INSERT INTO banners (content, status) VALUES ($1, $2) RETURNING id;"

	var bannerId int
	err = tx.QueryRow(query, BannerContent, banner.IsActive).Scan(&bannerId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if banner.FeatureID != "" {
		query := "INSERT INTO banner_features (banner_id, feature_id) VALUES ($1, $2);"
		_, err := tx.Exec(query, bannerId, banner.FeatureID)
		if err != nil {
			return fmt.Errorf("%s: failed to insert feature ID: %w", op, err)
		}
	}

	if len(banner.TagIds) != 0 {
		query := "INSERT INTO banner_tags (banner_id, tag_id) VALUES ($1, $2);"
		for _, tagId := range banner.TagIds {
			_, err = tx.Exec(query, bannerId, tagId)
			if err != nil {
				return fmt.Errorf("%s: failed to insert tag IDs: %w", op, err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}
