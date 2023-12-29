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
	RoomsKey    = "rooms"
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

func (s *RedisStorage) ListNodes(_ context.Context, region string) ([]*core.Node, error) {
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

func (s *RedisStorage) DeleteNode(_ context.Context, region string, nodeID string) error {
	key := NodesPrefix + region
	return s.rc.HDel(s.ctx, key, nodeID).Err()
}

func (s *RedisStorage) StoreRoom(_ context.Context, room *core.Room) error {
	data, err := proto.Marshal(room)
	if err != nil {
		return err
	}

	return s.rc.HSet(s.ctx, RoomsKey, room.Id, data).Err()
}

func (s *RedisStorage) LoadRoom(_ context.Context, roomID string) (*core.Room, error) {
	data, err := s.rc.HGet(s.ctx, RoomsKey, roomID).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, RoomNotFoundErr
		}
		return nil, err
	}

	var room core.Room
	err = proto.Unmarshal([]byte(data), &room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (s *RedisStorage) DeleteRoom(_ context.Context, roomID string) error {
	return s.rc.HDel(s.ctx, RoomsKey, roomID).Err()
}

func (s *RedisStorage) ListRooms(_ context.Context) ([]*core.Room, error) {
	items, err := s.rc.HVals(s.ctx, RoomsKey).Result()
	if err != nil {
		return nil, err
	}

	rooms := make([]*core.Room, 0)

	for _, n := range items {
		var room core.Room

		err := proto.Unmarshal([]byte(n), &room)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}

	return rooms, nil
}
