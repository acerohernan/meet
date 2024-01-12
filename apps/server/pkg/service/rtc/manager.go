package rtc

import (
	"context"
	"sync"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/storage"
)

var (
	// 1hr in seconds
	DefaultRoomMaxDuration = 3600
)

type rtcManager struct {
	ctx    context.Context
	mu     sync.Mutex
	store  storage.ObjectStore
	router *router.Router

	// map of roomID -> Room
	rooms map[string]*Room
	// map of participantID -> Participant
	participants map[string]*Participant
}

func NewRTCManager(store storage.ObjectStore, router *router.Router) RTCManager {
	m := &rtcManager{
		mu:           sync.Mutex{},
		rooms:        make(map[string]*Room),
		participants: make(map[string]*Participant),
		store:        store,
		ctx:          context.Background(),
		router:       router,
	}

	m.router.OnNodeMessage(m.handleNodeMessage)

	return m
}

func (m *rtcManager) CreateRoom(ctx context.Context, roomID string) error {
	// create room entity
	room := NewRoom(&core.Room{
		Id:              roomID,
		StartedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
		NumParticipants: 0,
		MaxDuration:     uint32(DefaultRoomMaxDuration),
	})

	// store locally
	m.mu.Lock()
	m.rooms[roomID] = room
	m.mu.Unlock()

	// store in object store
	if err := m.store.StoreRoom(ctx, room.proto); err != nil {
		return err
	}

	return nil
}

func (m *rtcManager) StartParticipantSignal(grants *auth.Grants) error {

	return nil
}
