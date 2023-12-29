package storage

import (
	"context"

	"github.com/acerohernan/meet/core"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . ObjectStore
type ObjectStore interface {
	StoreNode(ctx context.Context, region string, node *core.Node) error
	ListNodes(ctx context.Context, region string) ([]*core.Node, error)
	DeleteNode(ctx context.Context, region string, nodeID string) error

	StoreRoom(ctx context.Context, room *core.Room) error
	LoadRoom(ctx context.Context, roomID string) (*core.Room, error)
	DeleteRoom(ctx context.Context, roomID string) error
	ListRooms(ctx context.Context) ([]*core.Room, error)
}
