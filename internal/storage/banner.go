package storage

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
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
		err := rows.Scan(&banner.BannerID, &banner.FeatureID, pq.Array(&banner.TagIDs), &contentBytes, &banner.Status, &banner.CreatedAt, &banner.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		err = json.Unmarshal(contentBytes, &banner.Content)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
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

	if banner.FeatureID != 0 {
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

func (p *Postgres) DeleteBanner(bannerID int64) error {
	const op = "postgres.DeleteBanner"

	tx, err := p.database.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	query := `
		DELETE FROM banner_features
		WHERE banner_id = $1
	`
	_, err = tx.Exec(query, bannerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query = `
		DELETE FROM banner_tags
		WHERE banner_id = $1
	`
	_, err = tx.Exec(query, bannerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query = `
		DELETE FROM banners
		WHERE id = $1
	`
	_, err = tx.Exec(query, bannerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (p *Postgres) ChangeBanner(bannerID int64, banner models.ChangeBannerReq) error {
	const op = "postgres.ChangeBanner"

	tx, err := p.database.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	if banner.NewBanner != nil {
		query := `UPDATE banners SET content = $1 WHERE id = $2`
		bannerContent, err := json.Marshal(banner.NewBanner)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		_, err = tx.Exec(query, bannerContent, bannerID)
	}

	// TODO: Разобраться как проверять значения и реализовать работу функции через три запроса к бд

	return nil
}
