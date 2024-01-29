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

	// map of guestID -> Guest
	pendingGuests map[string]*core.Guest
}

func NewRoom(proto *core.Room) *Room {
	return &Room{
		proto:         proto,
		participants:  make(map[string]*Participant),
		pendingGuests: make(map[string]*core.Guest),
		mu:            sync.RWMutex{},
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

func (r *Room) ParticipantsMap() map[string]*Participant {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.participants
}

func (r *Room) Participants() []*Participant {
	r.mu.RLock()
	participants := r.participants
	r.mu.RUnlock()

	pList := make([]*Participant, 0)
	for _, p := range participants {
		pList = append(pList, p)
	}

	return pList
}

func (r *Room) ParticipantsProto() []*core.Participant {
	r.mu.RLock()
	participants := r.participants
	r.mu.RUnlock()

	pList := make([]*core.Participant, 0)
	for _, p := range participants {
		pList = append(pList, p.ToProto())
	}

	return pList
}

func (r *Room) GuestsProto() []*core.Guest {
	r.mu.RLock()
	guests := r.pendingGuests
	r.mu.RUnlock()

	gList := make([]*core.Guest, 0)
	for _, g := range guests {
		gList = append(gList, g)
	}

	return gList
}

func (r *Room) NumParticipants() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.participants)
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

func (r *Room) AddGuest(guest *core.Guest) {
	r.mu.Lock()
	r.pendingGuests[guest.Id] = guest
	r.mu.Unlock()
}

func (r *Room) GetGuest(guestId string) *core.Guest {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.pendingGuests[guestId]
}

func (r *Room) DeleteGuest(guestID string) {
	r.mu.Lock()
	delete(r.pendingGuests, guestID)
	r.mu.Unlock()
}
