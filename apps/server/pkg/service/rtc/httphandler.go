package rtc

import (
	"net/http"

	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (m *rtcManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	grants, err := auth.GetGrantsFromCTX(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		logger.Errow("error at upgrade ws connection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// start participant signal
	err = m.StartParticipantSignal(conn, grants)

	if err != nil {
		logger.Errow("couldnt start participant signal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO:
	// store participant in redis with node id
	// get node for roomID
	// send node message to start session
	// add participant to local store
	// write node message to node with active connection to participant (SignalRequest)
	// read the message and send to participant (SignalResponse)
}
