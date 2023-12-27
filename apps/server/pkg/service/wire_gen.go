// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/acerohernan/meet/pkg/config"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config) (*Server, error) {
	roomService := NewRoomService()
	server := NewServer(conf, roomService)
	return server, nil
}
