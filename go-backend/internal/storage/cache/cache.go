package cache

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var Cache *redis.Client = nil

func Connect() {
	rdUrl := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(rdUrl)
	if err != nil {
		panic(err)
	}

	Cache = redis.NewClient(opt)
}
