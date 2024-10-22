package redis

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "fullstack-lms-go/internal/config"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
        Password: cfg.RedisPassword,
        DB:       cfg.RedisDB,
    })

    // Test the connection
    ctx := context.Background()
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Redis: %w", err)
    }

    return rdb, nil
}