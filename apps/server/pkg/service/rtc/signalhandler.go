package rtc

import (
	"github.com/acerohernan/meet/core"
)

func (m *rtcManager) handleSignalRequest(room *Room, participant *Participant, req *core.SignalRequest) error {
	switch msg := req.GetRequest().(type) {
	case *core.SignalRequest_StartSession:
		// create participant
		p := NewParticipant(&core.Participant{
			Id:          msg.StartSession.Id,
			Name:        msg.StartSession.Name,
			Permissions: msg.StartSession.Permissions,
		}, m.router, req.GetStartSession().NodeId, room.ID())

		// store participant in room
		room.AddParticipant(p)

		// send join response
		if err := p.SendJoinResponse(&core.JoinResponse{
			Room: room.ToProto(),
		}); err != nil {
			return err
		}
	case *core.SignalRequest_CloseSession:
		room.DeleteParticipant(req.ParticipantId)
	}

	return nil
}
