package router

import (
	"sync"

	"github.com/acerohernan/meet/core"
)

type LocalMessenger struct {
	mu      sync.RWMutex
	msgChan chan *core.NodeMessage
}

func NewLocalMessenger() *LocalMessenger {
	return &LocalMessenger{
		mu:      sync.RWMutex{},
		msgChan: make(chan *core.NodeMessage),
	}
}

func (m *LocalMessenger) WriteMessage(_ string, msg *core.NodeMessage) error {
	m.mu.Lock()
	m.msgChan <- msg
	m.mu.Unlock()
	return nil
}

func (m *LocalMessenger) ReadChan() <-chan *core.NodeMessage {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.msgChan
}
