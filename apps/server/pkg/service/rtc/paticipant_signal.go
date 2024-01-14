package rtc

import (
	"github.com/acerohernan/meet/core"
)

func (p *Participant) SendJoinResponse(joinResponse *core.JoinResponse) error {
	return p.writeResponse(&core.SignalResponse{
		ParticipantId: p.ID(),
		RoomId:        p.RoomID(),
		Response: &core.SignalResponse_JoinResponse{
			JoinResponse: joinResponse,
		}})
}

func (p *Participant) SendRefreshToken(token string) error {
	return p.writeResponse(&core.SignalResponse{
		ParticipantId: p.ID(),
		RoomId:        p.RoomID(),
		Response: &core.SignalResponse_RefreshToken{
			RefreshToken: &core.RefreshToken{
				Token: token,
			},
		}})
}

func (p *Participant) writeResponse(res *core.SignalResponse) error {
	return p.router.SendNodeMessage(p.NodeID(), &core.NodeMessage{
		Message: &core.NodeMessage_SignalResponse{
			SignalResponse: res,
		}})
}
