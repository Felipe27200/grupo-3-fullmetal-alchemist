package queue

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() {
	addr := os.Getenv("REDIS_HOST")
	if addr == "" {
		addr = "alchemy-redis:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
}
