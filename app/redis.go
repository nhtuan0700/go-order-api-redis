package app

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(config Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
	})
}

func StartRedis(ctx context.Context, rdb *redis.Client) error {
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("fail to connect redis %w", err)
	}
	log.Println("Connect redis successfully")

	return nil
}

func StopRedis(rdb *redis.Client) {
	if err := rdb.Close(); err != nil {
		fmt.Println("Failed to close redis", err)
	}
}
