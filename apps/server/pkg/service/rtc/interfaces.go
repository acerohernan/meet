package rtc

import (
	"context"
	"net/http"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/service/auth"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RTCManager
type RTCManager interface {
	ServeWS(w http.ResponseWriter, r *http.Request)
	ServeJoinRequest(w http.ResponseWriter, r *http.Request)

	CreateRoom(ctx context.Context, roomID string) error
	GetRoom(ctx context.Context, roomID string) (*Room, error)
	StartParticipantSignal(nodeID string, grants *auth.Grants) (*core.SignalResponse, error)
	Close() error
}
