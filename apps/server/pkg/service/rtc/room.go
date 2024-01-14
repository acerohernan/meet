package rtc

import (
	"sync"

	"github.com/acerohernan/meet/core"
)

type Room struct {
	mu    sync.RWMutex
	proto *core.Room

	// map of participantID -> Participant
	participants map[string]*Participant
}

func NewRoom(proto *core.Room) *Room {
	return &Room{
		proto:        proto,
		participants: make(map[string]*Participant),
		mu:           sync.RWMutex{},
	}
}

func (r *Room) ID() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.proto.Id
}

func (r *Room) ToProto() *core.Room {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.proto
}

func (r *Room) AddParticipant(p *Participant) {
	r.mu.Lock()
	r.participants[p.ID()] = p
	r.proto.NumParticipants += 1
	r.mu.Unlock()
}

func (r *Room) GetParticipant(participantID string) *Participant {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.participants[participantID]
}

func (r *Room) DeleteParticipant(participantID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.participants[participantID] == nil {
		return
	}

	delete(r.participants, participantID)
	r.proto.NumParticipants -= 1
}
