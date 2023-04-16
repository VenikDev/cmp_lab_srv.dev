package redis

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/model/favorite"
	"context"
	"errors"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"time"
)

var (
	// redis client
	RedisClient *redis.Client
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

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),     // Redis server address and port
		Password: os.Getenv("REDIS_PASSWORD"), // Redis server password, if any
		DB:       dbNumber,                    // Redis database number to use (0-15)
	})

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil && RedisClient == nil {
		clog.Logger.Fatal("Redis...", "Can't connect to redis...", "FAIL")
	} else {
		clog.Logger.Info("Redis...", "Connected to Redis", "OK")
	}
}

// AddToPopular
// This code adds a key-value pair to a Redis database.
// The key is given as an argument to the function `AddToPopular`. If the key is empty,
// the function will return an error with the message "key is empty".
// The edited key is created by adding a prefix `RKW_POPULAR` to the given key. Then,
// the function checks if the value exists in the Redis database using the `Get` method. If the value does not exist,
// the function creates a new key-value pair with the given key and a value of 1 that expires after 24 hours.
// If the value already exists, the function increments the value associated with the key by 1 using a pipeline,
// and logs the edited key and the new value. Finally,
// the function returns nil if successful or an error if there was an issue executing the Redis commands.
func AddToPopular(key string) error {
	if key == "" {
		return errors.New("key is empty")
	}

	editedKey := RKW_POPULAR + key
	// check if value is not exists
	if valueOfKey := RedisClient.Get(ctx, editedKey); valueOfKey.Err() != nil {
		// save new value in redis on one day
		RedisClient.Set(ctx, editedKey, 1, time.Hour*24)

		clog.Logger.Info("Redis", "create value", editedKey)
	} else {
		pipe := RedisClient.Pipeline()
		incr := pipe.Incr(ctx, editedKey)

		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}

		clog.Logger.Info("Redis", editedKey, incr.Val())
	}

	return nil
}

// GetFavorite
// This code defines a function called `GetFavorite` that returns a slice of `favorite.
// Favorite` structs and an error. The function scans a Redis database using a wildcard key pattern,
// iterating over all keys that match the pattern. For each matching key, it retrieves the corresponding value,
// which is assumed to be an integer count. It then creates a new `favorite.
// Favorite` struct with the key name as the `Name` field and the count as the `Count` field,
// and appends it to the result slice. Finally,
// it returns the result slice and any errors encountered during scanning or parsing.
func GetFavorite() ([]favorite.Favorite, error) {
	var result []favorite.Favorite
	// key for parsing
	keyWord := RKW_POPULAR + "*"
	// get iterator
	iter := RedisClient.Scan(ctx, 0, keyWord, 0).Iterator()

	// send error
	if iter.Err() != nil {
		return nil, iter.Err()
	}

	// for each of keyword
	for iter.Next(ctx) {
		if getter := RedisClient.Get(ctx, iter.Val()); getter.Err() == nil {
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

// GetAnalysisByCity
func GetAnalysisByCity(city string) (string, error) {
	jsonData, err := RedisClient.Get(ctx, city).Result()
	if err != nil {
		return "", err
	}
	return jsonData, nil
}

// AddAnalysisByCity
func AddAnalysisByCity(city string, analysis model.LabAndListAnalyses) error {
	jsonData, err := json.Marshal(analysis)
	if err != nil {
		return err
	}

	errSet := RedisClient.Set(ctx, city, jsonData, time.Hour*24)
	if errSet != nil {
		return err
	}

	return nil
}
