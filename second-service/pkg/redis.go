package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(db int, ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis_container" + ":" + "6379",
		Password: "",
		DB:       db,
	})

	err := client.Ping(ctx).Err()
	return client, err
}
