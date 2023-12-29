package router

import "github.com/acerohernan/meet/core"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Monitor
type Monitor interface {
	GetCpuUsage() (float32, error)
	GetMemUsage() (float32, error)
}

//counterfeiter:generate . Messenger
type Messenger interface {
	WriteMessage(nodeID string, msg *core.NodeMessage) error
	ReadChan() <-chan *core.NodeMessage
}
