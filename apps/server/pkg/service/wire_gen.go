// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/rtc"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
	"github.com/redis/go-redis/v9"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config, localNode *core.Node) (*Server, error) {
	authService := auth.NewAuthService(conf)
	authMiddleware := auth.NewAuthMiddleware(authService)
	universalClient := getRedisClient(conf)
	objectStore := getObjectStore(universalClient)
	monitor, err := router.NewMonitor()
	if err != nil {
		return nil, err
	}
	messenger := getMessenger(universalClient, localNode)
	routerRouter := router.NewRouter(conf, localNode, objectStore, monitor, messenger)
	roomService := NewRoomService(routerRouter, objectStore, authService)
	rtcManager := rtc.NewRTCManager(objectStore, routerRouter)
	server := NewServer(conf, authMiddleware, roomService, routerRouter, rtcManager)
	return server, nil
}

// wire.go:

func getRedisClient(conf *config.Config) redis.UniversalClient {
	rc, err := utils.CreateRedisClient(conf.Redis)

	if err != nil {
		logger.Infow("error at creating redis client")
		return nil
	}

	return rc
}

func getObjectStore(rc redis.UniversalClient) storage.ObjectStore {
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
