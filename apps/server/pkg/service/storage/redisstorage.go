package storage

import (
	"context"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

const (
	NodesKey = "nodes"
)

type RedisStorage struct {
	rc  redis.UniversalClient
	ctx context.Context
}

func NewRedisStorage(rc redis.UniversalClient) *RedisStorage {
	logger.Infow("using redis storage as in-memory storage")
	return &RedisStorage{
		rc:  rc,
		ctx: context.Background(),
	}
}

func (s *RedisStorage) StoreNode(_ context.Context, node *core.Node) error {
	data, err := proto.Marshal(node)

	if err != nil {
		return err
	}

	return s.rc.HSet(s.ctx, NodesKey, node.Id, data).Err()
}

func (s *RedisStorage) ListNodes(ctx context.Context) ([]*core.Node, error) {
	items, err := s.rc.HVals(s.ctx, NodesKey).Result()

	if err != nil {
		return nil, err
	}

	nodes := make([]*core.Node, 0)

	for _, n := range items {
		var node core.Node

		err := proto.Unmarshal([]byte(n), &node)

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, &node)
	}

	return nodes, nil
}

func (s *RedisStorage) DeleteNode(ctx context.Context, nodeID string) error {
	return s.rc.HDel(s.ctx, NodesKey, nodeID).Err()
}
