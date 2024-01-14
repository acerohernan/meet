package rtc

import (
	"context"
	"errors"
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
	ctx       context.Context
	mu        sync.RWMutex
	store     storage.ObjectStore
	router    *router.Router
	localNode *core.Node

	// map of participantID -> SignalResponse
	participantResponses map[string]chan *core.SignalResponse

	// map of roomID -> Room
	rooms map[string]*Room
}

func NewRTCManager(localNode *core.Node, store storage.ObjectStore, router *router.Router) RTCManager {
	m := &rtcManager{
		mu:                   sync.RWMutex{},
		store:                store,
		ctx:                  context.Background(),
		router:               router,
		localNode:            localNode,
		rooms:                make(map[string]*Room),
		participantResponses: make(map[string]chan *core.SignalResponse),
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

	m.mu.RLock()
	nodeID := m.localNode.Id
	m.mu.RUnlock()

	// store room in local node
	if err := m.store.SetNodeForRoom(ctx, roomID, nodeID); err != nil {
		return err
	}

	return nil
}

func (m *rtcManager) GetRoom(_ context.Context, roomID string) (*Room, error) {
	m.mu.RLock()
	room := m.rooms[roomID]
	m.mu.RUnlock()

	if room == nil {
		return nil, RoomNotExistsInLocalNode
	}
	return room, nil
}

func (m *rtcManager) StartParticipantSignal(nodeID string, grants *auth.Grants) (*core.SignalResponse, error) {
	msg := &core.NodeMessage{
		Message: &core.NodeMessage_SignalRequest{
			SignalRequest: &core.SignalRequest{
				ParticipantId: grants.ID,
				RoomId:        grants.RoomID,
				Request: &core.SignalRequest_StartSession{
					StartSession: &core.StartSession{
						Id:     grants.ID,
						Name:   grants.Name,
						NodeId: m.localNode.Id,
						Permissions: &core.ParticipantPermissions{
							RoomAdmin: grants.RoomAdmin,
						},
					},
				},
			},
		},
	}

	// create chann for listening participant response
	m.mu.Lock()
	m.participantResponses[grants.ID] = make(chan *core.SignalResponse)
	m.mu.Unlock()

	if err := m.router.SendNodeMessage(nodeID, msg); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// delete created chann
			m.mu.Lock()
			delete(m.participantResponses, grants.ID)
			m.mu.Unlock()
			return nil, errors.New("timeout at waiting for participant signal")
		case msg := <-m.GetParticipantResponses(grants.ID):
			if msg == nil {
				return nil, errors.New("signal response not sended by host node")
			}
			return msg, nil
		}
	}
}

func (m *rtcManager) CloseParticipantSignal(nodeID string, grants *auth.Grants) error {
	msg := &core.NodeMessage{
		Message: &core.NodeMessage_SignalRequest{
			SignalRequest: &core.SignalRequest{
				ParticipantId: grants.ID,
				RoomId:        grants.RoomID,
				Request:       &core.SignalRequest_CloseSession{},
			},
		},
	}

	return m.router.SendNodeMessage(nodeID, msg)
}

func (m *rtcManager) ReceiveParticipantResponse(participantID string, res *core.SignalResponse) {
	m.mu.RLock()
	chann := m.participantResponses[res.ParticipantId]
	m.mu.RUnlock()
	// if channel do not exists, ignore message
	if chann == nil {
		return
	}
	chann <- res
}

func (m *rtcManager) GetParticipantResponses(participantID string) chan *core.SignalResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.participantResponses[participantID]
}
