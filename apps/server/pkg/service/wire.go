//go:build wireinject
// +build wireinject

package service

import (
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/storage"
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
	return storage.NewLocalStorage()
}

/*
func getRouter(storage storage.InMemoryStorage) *router.Router {
	return router.NewRouter(storage)
} */
