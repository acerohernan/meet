package rtc

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

func (r *Room) SendParticipantConnected(participant *core.Participant) {
	for _, p := range r.ParticipantsMap() {
		// omit new connected participant
		if p.ID() == participant.Id {
			continue
		}

		if err := p.SendParticipantConnected(participant); err != nil {
			logger.Infow("error at sending participant connected", "participant", p.ID())
		}
	}
}

func (r *Room) SendParticipantDisconnected(participantID string) {
	for _, p := range r.ParticipantsMap() {
		// omit disconnected participant
		if p.ID() == participantID {
			continue
		}

		if err := p.SendParticipantDisconnected(participantID); err != nil {
			logger.Infow("error at sending participant connected", "participant", p.ID())
		}
	}
}

func (r *Room) SendNewGuestRequest(guest *core.Guest) {
	for _, p := range r.ParticipantsMap() {
		// only notify room admins
		if !p.Permissions().RoomAdmin {
			continue
		}

		if err := p.SendNewGuestRequest(guest); err != nil {
			logger.Infow("error at sending new guest request", "participant", p.ID())
		}
	}
}

func (r *Room) SendGuestRequestCancelled(guestID string) {
	for _, p := range r.ParticipantsMap() {
		// only notify room admins
		if !p.Permissions().RoomAdmin {
			continue
		}

		if err := p.SendGuestRequestCancelled(guestID); err != nil {
			logger.Infow("error at sending guest request cancelled", "participant", p.ID())
		}
	}
}
