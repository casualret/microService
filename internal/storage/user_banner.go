package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"urlshortener/internal/models"
)

func (p *Postgres) GetUserBanner(req models.GetUserBannerReq) (*models.BannerWithDetails, error) {
	const op = "postgres.GetUserBanner"
	query := `
		SELECT b.id AS banner_id, bf.feature_id, array_agg(bt.tag_id) AS tag_ids, b.content, b.status, b.created_at, b.updated_at
		from banners b
		LEFT JOIN banner_features bf ON b.id = bf.banner_id
		LEFT JOIN banner_tags bt ON b.id = bt.banner_id
		WHERE bf.feature_id = $1
		GROUP BY b.id, bf.feature_id
		HAVING $2 = ANY(array_agg(bt.tag_id));
	`

	var banner models.BannerWithDetails
	var contentBytes []byte

	err := p.database.QueryRow(query, req.FeatureID, req.TagID).Scan(&banner.BannerID, &banner.FeatureID,
		pq.Array(&banner.TagIDs), &contentBytes, &banner.Status, &banner.CreatedAt, &banner.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(contentBytes, &banner.Content)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &banner, nil
}
