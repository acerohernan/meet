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

	GetNodeForRoom(ctx context.Context, roomID string) (string, error)
	SetNodeForRoom(ctx context.Context, roomID string, nodeID string) error
	DeleteNodeForRoom(ctx context.Context, roomID string) error

	StoreParticipant(ctx context.Context, roomID string, participant *core.Participant) error
	LoadParticipant(ctx context.Context, roomID string, participantID string) (*core.Participant, error)
	DeleteParticipant(ctx context.Context, roomID string, participantID string) error
	ListParticipants(ctx context.Context, roomID string) ([]*core.Participant, error)

	GetNodeForParticipant(ctx context.Context, roomID string) (string, error)
	SetNodeForParticipant(ctx context.Context, participantID string, nodeID string) error
	DeleteNodeForParticipant(ctx context.Context, participantID string) error
}
