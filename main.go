package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/storage"
)

func main() {

	//setup logger

	var logger *slog.Logger

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	//load config

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Error("error reading config", err)
		panic(err)
	}

	//connect database && cash

	db, err := storage.MustNewStorage(cfg)
	if err != nil {
		logger.Error("error setup storage", err)
		panic(err)
	}

	//init service

	service, err := NewService(db, logger)
	_ = service

	//init handlers && routs

	handlers := InitRoutes()

	//start server

	err = handlers.Run("localhost:8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Operation interface {
	GetBanners(ctx context.Context, src UserBanner) (*models.BannerWithDetails, error)
}

type Service struct {
	Operations
}

func NewService(db *storage.Postgres, logger *slog.Logger) (*Service, error) {
	return &Service{Operations: NewBannerMana}, nil
}

func InitRoutes() *gin.Engine {

	//r := gin.Default()
	r := gin.New()

	r.GET("/user_banner", func(c *gin.Context) {
		c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	})

	return r
}
