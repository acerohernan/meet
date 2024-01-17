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
