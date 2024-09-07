package services

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func SetupRedis() (*redis.Client, error) {
	redisClient, err := createRedisConnection()

	if err != nil {
		return nil, err
	}

	return redisClient, err

}

func createRedisConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	// Test the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
}

var redisClient, err = SetupRedis()
