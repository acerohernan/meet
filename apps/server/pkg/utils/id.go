package utils

import (
	"github.com/lithammer/shortuuid/v4"
)

const (
	idSize = 12
)

var (
	NodePrefix        = "NO_"
	RoomPrefix        = "RO_"
	ParticipantPrefix = "PA_"
	GuestPrefix       = "GE_"
)

func NewId(prefix string) string {
	return prefix + shortuuid.New()[:idSize]
}
