package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Redis client configuration to connect through HAProxy
	client := redis.NewClient(&redis.Options{
		Addr: "haproxy:6380",
	})

	// Testing Redis connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)

	// Setting a key in Redis
	err = client.Set(ctx, "example_key", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// Getting the value of the key from Redis
	val, err := client.Get(ctx, "example_key").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Println("example_key:", val)
}
