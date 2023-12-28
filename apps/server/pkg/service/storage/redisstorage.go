package storage

import (
	"context"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

const (
	NodesPrefix = "nodes:"
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

func (s *RedisStorage) StoreNode(_ context.Context, region string, node *core.Node) error {
	key := NodesPrefix + region

	data, err := proto.Marshal(node)

	if err != nil {
		return err
	}

	return s.rc.HSet(s.ctx, key, node.Id, data).Err()
}

func (s *RedisStorage) ListNodes(ctx context.Context, region string) ([]*core.Node, error) {
	key := NodesPrefix + region

	items, err := s.rc.HVals(s.ctx, key).Result()

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

func (s *RedisStorage) DeleteNode(ctx context.Context, region string, nodeID string) error {
	key := NodesPrefix + region
	return s.rc.HDel(s.ctx, key, nodeID).Err()
}
