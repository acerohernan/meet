//go:build wireinject
// +build wireinject

package service

import (
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/google/wire"
)

func InitializeServer(conf *config.Config) (*Server, error) {
	wire.Build(
		auth.NewAuthService,
		NewRoomService,
		NewServer,
	)

	return &Server{}, nil
}
