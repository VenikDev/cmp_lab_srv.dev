package redis

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/model"
	"cmp_lab/src/model/favorite"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"time"
)

var (
	// redis client
	RedisClient    *redis.Client
	ctx            = context.Background()
	ConnectSuccess = false
)

func TryConnectToRedis(host, password string, dbNumber int) (val string, error error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,     // Redis server address and port
		Password: password, // Redis server password, if any
		DB:       dbNumber, // Redis database number to use (0-15)
	})

	return RedisClient.Ping(context.Background()).Result()
}

func connectToRedisIsSuccess(err error) bool {
	if err != nil {
		clog.Error("[init/redis]", "Can't connect to redis...", err.Error())
		return false
	}

	clog.Info("[init/redis]", "Connected to Redis", "OK")
	return true
}

// InitRedis
// The code initializes a Redis client using the `go-redis/redis` package. It sets the Redis server address and port,
// password (if any), and database number to use.
// Then it sends a Ping command to the server to verify the connection. If the connection is successful,
// it logs an info message saying "Connected to Redis". Otherwise,
// it logs an error message saying "Can't connect to redis...".
func InitRedis() {
	dbNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		clog.Error("[init/redis]", "No parse number dbNumber", "OK")
		dbNumber = 0
	}

	redisHost := os.Getenv("REDIS_HOST")
	clog.Info("[init/redis]", "REDIS_HOST", redisHost)

	redisPassword := os.Getenv("REDIS_PASSWORD")
	clog.Info("[init/redis]", "REDIS_PASSWORD", redisPassword)

	_, err = TryConnectToRedis(redisHost, redisPassword, dbNumber)
	if !connectToRedisIsSuccess(err) {
		// for testing
		host := "http://localhost:6379"
		_, err = TryConnectToRedis(host, os.Getenv("REDIS_PASSWORD"), dbNumber)
		if err == nil {
			ConnectSuccess = true
		}
	} else {
		ConnectSuccess = true
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
		oneDay := time.Hour * 24
		err := RedisClient.Set(ctx, editedKey, 1, oneDay).Err()
		if err != nil {
			clog.Info("[add/redis]", "Error adding popular", err)
		} else {
			clog.Info("[add/redis]", "create value", editedKey)
		}
	} else {
		err := RedisClient.Incr(ctx, editedKey).Err()
		if err != nil {
			clog.Info("[add/redis]", "Error increment popular", err)
		}
	}

	return nil
}

// GetPopular
// This code defines a function called `GetPopular` that returns a slice of `favorite.
// Favorite` structs and an error. The function scans a Redis database using a wildcard key pattern,
// iterating over all keys that match the pattern. For each matching key, it retrieves the corresponding value,
// which is assumed to be an integer count. It then creates a new `favorite.
// Favorite` struct with the key name as the `Name` field and the count as the `Count` field,
// and appends it to the result slice. Finally,
// it returns the result slice and any errors encountered during scanning or parsing.
func GetPopular() ([]favorite.Favorite, error) {
	// pre-allocate a slice with a sufficient capacity
	result := make([]favorite.Favorite, 0, 100)

	// key for parsing
	keyWord := RKW_POPULAR + "*"

	// get iterator
	iter := RedisClient.Scan(ctx, 0, keyWord, 0).Iterator()

	// send error
	if err := iter.Err(); err != nil {
		return nil, err
	}

	// for each of keyword
	for iter.Next(ctx) {
		val := iter.Val()
		name := val[len(keyWord)-1:]

		// parse
		getter := RedisClient.Get(ctx, val)
		if getter.Err() == nil {
			// add
			if count, err := getter.Int64(); err == nil {
				result = append(result, favorite.Favorite{
					Name:  name,
					Count: count,
				})
			}
		}
	}

	return result, nil
}

func GetAnalysisByCity(city string) (string, error) {
	jsonData, err := RedisClient.Get(ctx, city).Result()
	if err != nil {
		return "", err
	}
	return jsonData, nil
}

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
