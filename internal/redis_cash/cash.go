package redis_cash

import (
	"github.com/redis/go-redis/v9"
)

type RedisCash struct {
	Client *redis.Client
}

func NewRedisClient() *RedisCash {
	return &RedisCash{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
