package redis

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	RedisClient redis.Client
}

func NewRedis() Redis {
	var client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "12345",
	})

	return Redis{
		RedisClient: *client,
	}
}
