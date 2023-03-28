package redis

import (
	"comparisonLaboratories/src/clog"
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var (
	// redis client
	redisClient *redis.Client
)

// InitRedis
// The code initializes a Redis client using the `go-redis/redis` package. It sets the Redis server address and port,
// password (if any), and database number to use.
// Then it sends a Ping command to the server to verify the connection. If the connection is successful,
// it logs an info message saying "Connected to Redis". Otherwise,
// it logs an error message saying "Can't connect to redis...".
func InitRedis() {
	dbNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		clog.Logger.Fatal("Redis...", "No parse number dbNumber", "OK")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),     // Redis server address and port
		Password: os.Getenv("REDIS_PASSWORD"), // Redis server password, if any
		DB:       dbNumber,                    // Redis database number to use (0-15)
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		clog.Logger.Fatal("Redis...", "Can't connect to redis...", "FAIL")
	} else {
		clog.Logger.Info("Redis...", "Connected to Redis", "OK")
	}
}
