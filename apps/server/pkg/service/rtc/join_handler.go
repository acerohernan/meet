package rtc

import (
	"fmt"
	"net/http"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (m *rtcManager) ServeJoinRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	query := r.URL.Query()
	roomID := query.Get("roomId")
	name := query.Get("name")

	// params validation
	if roomID == "" || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start server sent events connection
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	logger.Infow("receiving request", "room", roomID, "name", name)

	nodeID, err := m.store.GetNodeForRoom(m.ctx, roomID)
	if err != nil {
		if err == storage.NodeForRoomNotFoundErr {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	m.mu.RLock()
	localNodeID := m.localNode.Id
	m.mu.RUnlock()

	guest := &core.Guest{
		Id:     utils.NewId(utils.GuestPrefix),
		Name:   name,
		NodeId: localNodeID,
	}

	msg := &core.NodeMessage{
		Message: &core.NodeMessage_GuestJoinRequest{
			GuestJoinRequest: &core.GuestJoinRequest{
				RoomId: roomID,
				Guest:  guest,
			},
		},
	}

	m.AddGuestJoinRequest(guest.Id)

	if err := m.router.SendNodeMessage(nodeID, msg); err != nil {
		m.DeleteGuestJoinRequest(guest.Id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		m.DeleteGuestJoinRequest(guest.Id)

		// notify request cancellation to room
		msg := &core.NodeMessage{
			Message: &core.NodeMessage_GuestRequestCancelled{
				GuestRequestCancelled: &core.GuestRequestCancelled{
					RoomId:  roomID,
					GuestId: guest.Id,
				},
			},
		}
		m.router.SendNodeMessage(nodeID, msg)
		return
	}()

	for {
		select {
		case msg := <-m.GetGuestResponseChan(guest.Id):
			jsonBytes, err := protojson.Marshal(msg)
			if err != nil {
				logger.Errorw("error at parsing message guest join answer", err)
				return
			}

			fmt.Fprintf(w, "data: %s\n\n", string(jsonBytes))
			w.(http.Flusher).Flush()
			return
		}
	}
}
