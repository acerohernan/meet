//go:build wireinject
// +build wireinject

package service

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitializeServer(conf *config.Config, localNode *core.Node) (*Server, error) {
	wire.Build(
		getRedisClient,
		getMessenger,
		getInMemoryStorage,
		router.NewMonitor,
		router.NewRouter,
		auth.NewAuthMiddleware,
		auth.NewAuthService,
		NewRoomService,
		NewServer,
	)

	return &Server{}, nil
}

func getRedisClient(conf *config.Config) redis.UniversalClient {
	rc, err := utils.CreateRedisClient(conf.Redis)

	if err != nil {
		logger.Infow("error at creating redis client")
		return nil
	}

	return rc
}

func getInMemoryStorage(rc redis.UniversalClient) storage.InMemoryStorage {
	if rc != nil {
		return storage.NewRedisStorage(rc)
	}
	return storage.NewLocalStorage()
}

func getMessenger(rc redis.UniversalClient, localNode *core.Node) router.Messenger {
	if rc != nil {
		return router.NewRedisMessenger(rc, localNode)
	}
	return router.NewLocalMessenger()
}
