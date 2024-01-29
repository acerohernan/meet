package rtc

import (
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/utils"
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
			Guests:       room.GuestsProto(),
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

	case *core.SignalRequest_AnswerGuestRequest:
		resChan := m.GetGuestResponseChan(msg.AnswerGuestRequest.GuestId)
		if resChan == nil {
			return nil
		}
		guest := room.GetGuest(msg.AnswerGuestRequest.GuestId)
		if guest == nil {
			return nil
		}

		res := &core.GuestJoinResponse{Guest: guest}

		switch msg.AnswerGuestRequest.Answer.(type) {
		case *core.AnswerGuestRequest_GuestDenied:
			res.Answer = &core.GuestJoinResponse_JoinDenied{}

		case *core.AnswerGuestRequest_GuestAccepted:
			// create new participant grants
			at, err := m.authSvc.NewAccessTokenFromGrants(&auth.Grants{
				ID:        utils.NewId(utils.ParticipantPrefix),
				Name:      guest.Name,
				RoomID:    room.ID(),
				RoomAdmin: false,
			})
			if err != nil {
				return err
			}

			res.Answer = &core.GuestJoinResponse_JoinApproved{
				JoinApproved: &core.JoinApproved{
					AccessToken: at.ToJWT(),
				},
			}

		default:
			return nil
		}

		resChan <- res
		room.DeleteGuest(msg.AnswerGuestRequest.GuestId)
	}

	return nil
}
