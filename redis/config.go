package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"dating-sim/dotenv"

	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func Init() {
	redisAddr := dotenv.Env("REDIS_ADDR")
	redisPassword := dotenv.Env("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(dotenv.Env("REDIS_DB"))
	if err != nil {
		panic("failed init redis: no redis DD/redis is not number")
	}

	// Create a Redis client
	Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Check the connection
	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Print("Connected to Redis successfully!")
	}
}

func Close() {
	if err := Client.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	} else {
		fmt.Println("Redis connection closed")
	}
}
