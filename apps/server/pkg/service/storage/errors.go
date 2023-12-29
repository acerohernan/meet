package storage

import "errors"

var (
	RoomNotFoundErr               = errors.New("room not found")
	ParticipantNotFoundErr        = errors.New("room not found")
	NodeForRoomNotFoundErr        = errors.New("node for room not found")
	NodeForParticipantNotFoundErr = errors.New("node for room not found")
)
