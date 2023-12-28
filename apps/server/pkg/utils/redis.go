package utils

import (
	"github.com/acerohernan/meet/pkg/config"
	"github.com/redis/go-redis/v9"
)

func CreateRedisClient(conf *config.RedisConfig) (rc redis.UniversalClient) {
	// if redis is not configured
	if conf.Address == "" {
		return nil
	}

	return redis.NewClient(&redis.Options{
		Addr: conf.Address,
	})
}
