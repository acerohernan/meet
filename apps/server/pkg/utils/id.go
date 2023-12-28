package utils

import (
	"github.com/lithammer/shortuuid/v4"
)

const (
	idSize = 12
)

var (
	NodePrefix = "NO_"
)

func NewId(prefix string) string {
	return prefix + shortuuid.New()[:idSize]
}
