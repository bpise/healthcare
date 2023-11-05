package cache

import (
	"context"
	"os"
	"time"

	"healthcare/logger"

	"github.com/go-redis/redis/v8"
)

const default_REDIS_URL = "redis://127.0.0.1:6379"

var redisClient *redis.Client

// InitRedis - Initialize Redis client.
func InitRedis(ctx context.Context) {
	if redisClient != nil {
		return
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = default_REDIS_URL
	}

	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	redisClient = redis.NewClient(opts)

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Infof("Connected to Redis ...")
}

// Closes the Redis client.
func Close() {
	_ = redisClient.Close()
}

// Set a key-value pair in Redis with an optional expiration time.
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redisClient.Set(ctx, key, value, expiration)
}

// Get the value associated with a key from Redis.
func Get(ctx context.Context, key string) *redis.StringCmd {
	return redisClient.Get(ctx, key)
}
