package rtc

import (
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
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
			Room:         room.ToProto(),
			Participants: room.ParticipantsProto(),
		}); err != nil {
			return err
		}

		// send new participant connected
		room.SendParticipantConnected(p.ToProto())

		// refresh access token inmmediately
		at, err := m.authSvc.NewAccessTokenFromGrants(p.Grants())
		if err != nil {
			return err
		}
		p.SendRefreshToken(at.ToJWT())

		// participant refresh token interval
		go func() {
			ticker := time.NewTicker(time.Minute * 5)
			for {
				select {
				case <-ticker.C:
					// verify that partipant is active before sending new token
					if room.GetParticipant(p.ID()) == nil {
						return
					}

					at, err := m.authSvc.NewAccessTokenFromGrants(p.Grants())
					if err != nil {
						continue
					}
					if err := p.SendRefreshToken(at.ToJWT()); err != nil {
						logger.Errorw("error at sending token refreshed to participant", err)
						continue
					}
				}
			}
		}()

	case *core.SignalRequest_CloseSession:
		room.DeleteParticipant(req.ParticipantId)
		room.SendParticipantDisconnected(req.ParticipantId)
	}

	return nil
}
