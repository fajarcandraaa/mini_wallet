package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func NewRedisConfig() (*redis.Client, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return nil, err
	}

	// Redis connection options
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	}

	// Create a new Redis client
	client := redis.NewClient(options)

	return client, nil
}
