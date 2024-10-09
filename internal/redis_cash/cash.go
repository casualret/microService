package redis_cash

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"microService/internal/models"
	"time"
)

type RedisCash struct {
	client *redis.Client
}

func NewRedisClient() *RedisCash {
	return &RedisCash{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (r *RedisCash) SaveBanner(ctx *context.Context, banners ...*models.BannerWithDetails) error {
	const op = "redis_cash.SaveBanner"

	bannerJSON, err := json.Marshal(banners)
	if err != nil {
		return err
	}
	for i := range banners {
		err = r.client.Set(*ctx, fmt.Sprintf("bannerID:%d", banners[i].BannerID), bannerJSON, 5*time.Minute).Err()
		if err != nil {
			log.Printf("%s: failed cash banner id = %d", op, banners[i].BannerID)
		}
	}
	return nil
}
