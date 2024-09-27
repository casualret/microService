package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"microService/internal/config"
)

type Postgres struct {
	database *sqlx.DB
}

func MustNewStorage(cfg *config.Config) (*Postgres, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresCfg.PgHost, cfg.PostgresCfg.PgPort, cfg.PostgresCfg.PgUser, cfg.PostgresCfg.PgPassword, cfg.PostgresCfg.PgDatabase, cfg.PostgresCfg.PgSslmode)
	db, err := sqlx.Connect("postgres", connInfo) // connect to postgres
	if err != nil {
		return nil, fmt.Errorf("postgres connect error: %v", err)
	}
	return &Postgres{database: db}, nil
}
