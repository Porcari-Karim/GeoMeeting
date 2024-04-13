package cache

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
)

var cache *redis.Client = nil

func Connect() {
	rdUrl := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(rdUrl)
	if err != nil {
		panic(err)
	}

	cache = redis.NewClient(opt)
}

func Close() {
	err := cache.Close()
	if err != nil {
		panic(err)
	}
}

func GetInstance() (*redis.Client, error) {
	if cache == nil {
		return nil, errors.New("error: Cache can't be 'nil' try to 'Initialize'")
	}
	return cache, nil
}
