package utils

import (
	"context"
	"errors"

	"github.com/acerohernan/meet/pkg/config"
	"github.com/redis/go-redis/v9"
)

func CreateRedisClient(conf *config.RedisConfig) (redis.UniversalClient, error) {
	// if redis is not configured
	if conf.Address == "" {
		return nil, errors.New("redis is not configured")
	}

	rc := redis.NewClient(&redis.Options{
		Addr: conf.Address,
	})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		return nil, errors.New("unable to connect to redis")
	}

	return rc, nil
}
