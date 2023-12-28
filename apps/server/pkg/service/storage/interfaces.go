package storage

import (
	"context"

	"github.com/acerohernan/meet/core"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . InMemoryStorage
type InMemoryStorage interface {
	StoreNode(ctx context.Context, node *core.Node) error
	ListNodes(ctx context.Context) ([]*core.Node, error)
	DeleteNode(ctx context.Context, nodeID string) error
}
