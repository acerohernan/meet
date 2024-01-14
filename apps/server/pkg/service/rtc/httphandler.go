package rtc

import (
	"errors"
	"net/http"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
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

	nodeID, err := m.store.GetNodeForRoom(m.ctx, grants.RoomID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := m.StartParticipantSignal(nodeID, grants)
	if err != nil {
		logger.Infow("error at start participant signal", err, "grants", grants)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		err := m.CloseParticipantSignal(nodeID, grants)
		if err != nil {
			logger.Errorw("error at closing participant signal", err)
		}
	}()

	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		logger.Errorw("error at upgrade ws connection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Infow("new ws connection established!", "participant", grants.ID)

	defer func() {
		logger.Infow("ws connection closed", "participant", grants.ID)
	}()

	// send join response
	data, err := proto.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go func() {
		for {
			select {
			case <-r.Context().Done():

				return
			case msg := <-m.GetParticipantResponses(grants.ID):
				logger.Infow("sending response to participant", "msg", msg)
				// send participant responses
				data, err := proto.Marshal(msg)
				if err != nil {
					logger.Errorw("error at encoding signal response", err, "participant", grants.ID)
				}

				err = conn.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					logger.Errorw("error at sending signal response", err, "participant", grants.ID)
				}
			}
		}

	}()

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(
				err,
				websocket.CloseAbnormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure,
				websocket.CloseNoStatusReceived,
			) {
				logger.Infow("exit ws read loop due to closed connection", "participant", grants.ID)
			} else {
				logger.Errorw("error reading from websocket", err)
			}
			return
		}

		msg := &core.SignalRequest{}

		switch messageType {
		case websocket.BinaryMessage:
			err := proto.Unmarshal(payload, msg)
			if err == nil {
				logger.Infow("signal request received", "msg", msg)
			}
		default:
			logger.Errorw("ws message not supported", errors.New("ws msg type not supported"), "type", messageType)
		}
	}
}
