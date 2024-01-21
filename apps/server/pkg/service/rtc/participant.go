package rtc

import (
	"sync"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
)

type Participant struct {
	mu     sync.RWMutex
	proto  *core.Participant
	nodeID string
	roomID string

	router *router.Router
}

func NewParticipant(proto *core.Participant, router *router.Router, nodeID string, roomID string) *Participant {
	return &Participant{
		router: router,

		proto:  proto,
		nodeID: nodeID,
		roomID: roomID,
	}
}

func (p *Participant) ID() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.proto.Id
}

func (p *Participant) ToProto() *core.Participant {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.proto
}

func (p *Participant) RoomID() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.roomID
}

func (p *Participant) NodeID() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.nodeID
}

func (p *Participant) Permissions() *core.ParticipantPermissions {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.proto.Permissions
}

func (p *Participant) Grants() *auth.Grants {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return &auth.Grants{
		ID:        p.proto.Id,
		Name:      p.proto.Name,
		RoomID:    p.roomID,
		RoomAdmin: p.proto.Permissions.RoomAdmin,
	}
}
