package rtc

import (
	"context"
	"net/http"

	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/gorilla/websocket"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . RTCManager
type RTCManager interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	CreateRoom(ctx context.Context, roomID string) error
	StartParticipantSignal(conn *websocket.Conn, grants *auth.Grants) error
}
