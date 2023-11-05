package cache

import (
	"context"
	"time"

	"healthcare/logger"

	"github.com/go-redis/redis/v8"
)

const default_REDIS_URL = "redis://127.0.0.1:6379"

var redisClient *redis.Client

// InitRedis - initialize Redis.
func InitRedis(ctx context.Context) {
	if redisClient != nil {
		return
	}

	opts, err := redis.ParseURL(default_REDIS_URL)
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

func Close() {
	_ = redisClient.Close()
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redisClient.Set(ctx, key, value, expiration)
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	return redisClient.Get(ctx, key)
}
