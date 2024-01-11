package rtc

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	// 1hr in seconds
	DefaultRoomMaxDuration = 3600
)

type rtcManager struct {
	ctx   context.Context
	mu    sync.Mutex
	store storage.ObjectStore
	// map of roomID -> Room
	rooms map[string]*Room
	// map of participantID -> Participant
	participants map[string]*Participant
}

func NewRTCManager(store storage.ObjectStore) RTCManager {
	return &rtcManager{
		mu:           sync.Mutex{},
		rooms:        make(map[string]*Room),
		participants: make(map[string]*Participant),
		store:        store,
		ctx:          context.Background(),
	}
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

func (m *rtcManager) StartParticipantSignal(conn *websocket.Conn, grants *auth.Grants) error {
	// create participant entity with function like sendResponse
	proto := &core.Participant{
		Id:   grants.ID,
		Name: grants.Name,
		Permissions: &core.ParticipantPermissions{
			RoomAdmin: grants.RoomAdmin,
		},
	}

	p := NewParticipant(conn, proto)

	// spin a go routine to read all signal request
	go m.signalWorker(conn, p)

	// save participant
	m.mu.Lock()
	m.participants[proto.Id] = p
	m.mu.Unlock()

	// store it
	if err := m.store.StoreParticipant(m.ctx, grants.RoomID, proto); err != nil {
		return err
	}

	return nil
}

func (m *rtcManager) signalWorker(conn *websocket.Conn, p *Participant) {
	// read all participant messages
	for {
		messageType, payload, err := conn.ReadMessage()

		if err != nil {
			logger.Errorw("error at reading ws message", err)
		}

		msg := &core.SignalRequest{}

		switch messageType {
		case websocket.BinaryMessage:
			err := proto.Unmarshal(payload, msg)
			if err == nil {
				logger.Infow("signal request received", "msg", msg)
			}
		default:
			logger.Errorw("ws message not supported", errors.New("ws msg type not supported"), "type", messageType)
		}

	}
}
