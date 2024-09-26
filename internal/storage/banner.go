package storage

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"time"
	"urlshortener/internal/models"
)

func (p *Postgres) GetBannersParams(req models.GetBannersReq) ([]*models.BannerWithDetails, error) {
	const op = "postgres.GetBannerParams"

	query := `
		SELECT b.id AS banner_id, bf.feature_id, array_agg(bt.tag_id) AS tag_ids, b.content, b.status, b.created_at, b.updated_at
		FROM banners b 
		LEFT JOIN banner_features bf ON b.id = bf.banner_id
		LEFT JOIN banner_tags bt ON b.id = bt.banner_id
	`

	args := make([]interface{}, 0)
	argIndex := 1
	if req.FeatureID != nil {
		query += fmt.Sprintf(" WHERE bf.feature_id = $%d", argIndex)
		args = append(args, *req.FeatureID)
		argIndex++
	}

	query += " GROUP BY b.id, bf.feature_id"

	if req.TagID != nil {
		//if req.FeatureID != nil {
		//	query += " AND"
		//} else {
		//	query += " WHERE"
		//}
		//query += fmt.Sprintf(" bt.tag_id = $%d", argIndex)
		query += fmt.Sprintf(" HAVING $%d = ANY(array_agg(bt.tag_id))", argIndex)
		args = append(args, *req.TagID)
		argIndex++
	}

	if req.Limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, *req.Limit)
		argIndex++
	}
	if req.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, *req.Offset)
	}

	query += ";"

	rows, err := p.database.Queryx(query, args...)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var bannerWD []*models.BannerWithDetails
	for rows.Next() {
		var banner models.BannerWithDetails
		var contentBytes []byte
		var tagIDs []int64
		var status bool
		var createdAt, updatedAt time.Time
		err := rows.Scan(&banner.BannerID, &banner.FeatureID, pq.Array(&tagIDs), &contentBytes, &status, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		err = json.Unmarshal(contentBytes, &banner.Content)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		banner.TagIDs = tagIDs
		banner.Status = &status
		banner.CreatedAt = &createdAt
		banner.UpdatedAt = &updatedAt
		bannerWD = append(bannerWD, &banner)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return bannerWD, nil
}

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
