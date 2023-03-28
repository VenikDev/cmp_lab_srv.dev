package redis

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/model/favorite"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"time"
)

var (
	// redis client
	redisClient *redis.Client
	ctx         = context.Background()
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

	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),     // Redis server address and port
		Password: os.Getenv("REDIS_PASSWORD"), // Redis server password, if any
		DB:       dbNumber,                    // Redis database number to use (0-15)
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil && redisClient == nil {
		clog.Logger.Fatal("Redis...", "Can't connect to redis...", "FAIL")
	} else {
		clog.Logger.Info("Redis...", "Connected to Redis", "OK")
	}
}

func AddKeyToRedis(key string) error {
	if key == "" {
		return errors.New("key is empty")
	}

	editedKey := RKW_FAVORITE + key
	// check if value is not exists
	if valueOfKey := redisClient.Get(ctx, editedKey); valueOfKey.Err() != nil {
		// save new value in redis on one day
		redisClient.Set(ctx, editedKey, 1, time.Hour*24)

		clog.Logger.Info("Redis", "create value", editedKey)
	} else {
		pipe := redisClient.Pipeline()
		incr := pipe.Incr(ctx, editedKey)

		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}

		clog.Logger.Info("Redis", editedKey, incr.Val())
	}

	return nil
}

func GetFavorite() ([]favorite.Favorite, error) {
	var result []favorite.Favorite
	// key for parsing
	keyWord := RKW_FAVORITE + "*"
	// get iterator
	iter := redisClient.Scan(ctx, 0, keyWord, 0).Iterator()

	// send error
	if iter.Err() != nil {
		return nil, iter.Err()
	}

	// for each of keyword
	for iter.Next(ctx) {
		if getter := redisClient.Get(ctx, iter.Val()); getter.Err() == nil {
			// parse
			if count, err := getter.Int64(); err == nil {
				// add
				result = append(result, favorite.Favorite{
					Name:  iter.Val()[len(keyWord)-1:],
					Count: count,
				})
			}
		}

	}
	return result, nil
}
