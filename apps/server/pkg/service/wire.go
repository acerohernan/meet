//go:build wireinject
// +build wireinject

package service

import (
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
	"github.com/google/wire"
)

func InitializeServer(conf *config.Config) (*Server, error) {
	wire.Build(
		getInMemoryStorage,
		router.NewRouter,
		auth.NewAuthMiddleware,
		auth.NewAuthService,
		NewRoomService,
		NewServer,
	)

	return &Server{}, nil
}

func getInMemoryStorage(conf *config.Config) storage.InMemoryStorage {
	rc := utils.CreateRedisClient(conf.Redis)

	if rc != nil {
		return storage.NewRedisStorage(rc)
	}

	return storage.NewLocalStorage()
}

/*
func getRouter(storage storage.InMemoryStorage) *router.Router {
	return router.NewRouter(storage)
} */
