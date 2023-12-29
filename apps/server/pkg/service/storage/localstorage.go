package storage

import (
	"context"
	"sync"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

type LocalStorage struct {
	mu sync.RWMutex
	// map of region -> map of nodeID -> Node
	nodes map[string]map[string]*core.Node
	// map of roomID -> Room
	rooms map[string]*core.Room
	// map of participantID -> Participant
	participants map[string]*core.Participant
	// map of roomID -> nodeID
	roomNodes map[string]string
	// map of participantID -> nodeID
	participantNodes map[string]string
}

func NewLocalStorage() *LocalStorage {
	logger.Infow("using local storage as in-memory storage")
	return &LocalStorage{
		mu:               sync.RWMutex{},
		nodes:            make(map[string]map[string]*core.Node),
		rooms:            make(map[string]*core.Room),
		roomNodes:        make(map[string]string),
		participantNodes: make(map[string]string),
	}
}

func (s *LocalStorage) StoreNode(ctx context.Context, region string, node *core.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.nodes[region] == nil {
		s.nodes[region] = make(map[string]*core.Node)
	}

	s.nodes[region][node.Id] = node

	return nil
}

func (s *LocalStorage) ListNodes(ctx context.Context, region string) ([]*core.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodes := s.nodes[region]

	if nodes == nil {
		return nil, nil
	}

	items := make([]*core.Node, 0)

	for _, n := range nodes {
		items = append(items, n)
	}

	return items, nil
}

func (s *LocalStorage) DeleteNode(ctx context.Context, region string, nodeID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.nodes[region] == nil {
		return nil
	}

	delete(s.nodes[region], nodeID)

	return nil
}

func (s *LocalStorage) StoreRoom(ctx context.Context, room *core.Room) error {
	s.mu.Lock()
	s.rooms[room.Id] = room
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) LoadRoom(ctx context.Context, roomID string) (*core.Room, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	room := s.rooms[roomID]
	if room == nil {
		return nil, RoomNotFoundErr
	}
	return room, nil
}

func (s *LocalStorage) DeleteRoom(ctx context.Context, roomID string) error {
	s.mu.Lock()
	delete(s.rooms, roomID)
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) ListRooms(ctx context.Context) ([]*core.Room, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rooms := make([]*core.Room, 0)
	for _, r := range s.rooms {
		rooms = append(rooms, r)
	}
	return rooms, nil
}

func (s *LocalStorage) GetNodeForRoom(ctx context.Context, roomID string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodeID := s.roomNodes[roomID]
	if nodeID == "" {
		return "", NodeForRoomNotFoundErr
	}
	return nodeID, nil
}

func (s *LocalStorage) SetNodeForRoom(ctx context.Context, roomID string, nodeID string) error {
	s.mu.Lock()
	s.roomNodes[roomID] = nodeID
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) DeleteNodeForRoom(ctx context.Context, roomID string) error {
	s.mu.Lock()
	delete(s.roomNodes, roomID)
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) StoreParticipant(ctx context.Context, roomID string, participant *core.Participant) error {
	s.mu.Lock()
	s.participants[participant.Id] = participant
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) LoadParticipant(ctx context.Context, roomID string, participantID string) (*core.Participant, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p := s.participants[participantID]
	if p == nil {
		return nil, ParticipantNotFoundErr
	}
	return p, nil
}

func (s *LocalStorage) DeleteParticipant(ctx context.Context, roomID string, participantID string) error {
	s.mu.Lock()
	delete(s.participants, participantID)
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) ListParticipants(ctx context.Context, roomID string) ([]*core.Participant, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	participants := make([]*core.Participant, 0)
	for _, r := range s.participants {
		participants = append(participants, r)
	}
	return participants, nil
}

func (s *LocalStorage) GetNodeForParticipant(ctx context.Context, participantID string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodeID := s.participantNodes[participantID]
	if nodeID == "" {
		return "", NodeForParticipantNotFoundErr
	}
	return nodeID, nil
}

func (s *LocalStorage) SetNodeForParticipant(ctx context.Context, participantID string, nodeID string) error {
	s.mu.Lock()
	s.participantNodes[participantID] = nodeID
	s.mu.Unlock()
	return nil
}

func (s *LocalStorage) DeleteNodeForParticipant(ctx context.Context, participantID string) error {
	s.mu.Lock()
	delete(s.participantNodes, participantID)
	s.mu.Unlock()
	return nil
}
