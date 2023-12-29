package rtc

import (
	"context"
	"sync"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/service/storage"
)

var (
	// 1hr in seconds
	DefaultRoomMaxDuration = 3600
)

type RoomManager struct {
	ctx   context.Context
	mu    sync.Mutex
	store storage.ObjectStore
	// map of roomID -> Room
	rooms map[string]*core.Room
}

func NewRoomManager(store storage.ObjectStore) *RoomManager {
	return &RoomManager{
		mu:    sync.Mutex{},
		rooms: make(map[string]*core.Room),
		store: store,
		ctx:   context.Background(),
	}
}

func (m *RoomManager) CreateRoom(ctx context.Context, roomID string) error {
	// create room entity
	room := &core.Room{
		Id:              roomID,
		StartedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
		NumParticipants: 0,
		MaxDuration:     uint32(DefaultRoomMaxDuration),
	}

	// store locally
	m.mu.Lock()
	m.rooms[roomID] = room
	m.mu.Unlock()

	// store in object store
	if err := m.store.StoreRoom(ctx, room); err != nil {
		return err
	}

	return nil
}
