package router

import (
	"context"
	"sync"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type RedisMessenger struct {
	ctx       context.Context
	mu        sync.RWMutex
	rc        redis.UniversalClient
	redisChan <-chan *redis.Message
	msgChan   chan *core.NodeMessage
}

func NewRedisMessenger(rc redis.UniversalClient, localNode *core.Node) *RedisMessenger {
	m := &RedisMessenger{
		rc:      rc,
		ctx:     context.Background(),
		msgChan: make(chan *core.NodeMessage),
		mu:      sync.RWMutex{},
	}

	m.redisChan = m.rc.Subscribe(m.ctx, localNode.Id).Channel()

	go m.messageProxy()

	return m
}

func (m *RedisMessenger) WriteMessage(nodeID string, msg *core.NodeMessage) error {
	data, err := proto.Marshal(msg)

	if err != nil {
		return err
	}

	return m.rc.Publish(m.ctx, nodeID, data).Err()
}

func (m *RedisMessenger) ReadChan() <-chan *core.NodeMessage {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.msgChan
}

func (m *RedisMessenger) messageProxy() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case msg := <-m.redisChan:
			var nodeMsg core.NodeMessage
			err := proto.Unmarshal([]byte(msg.Payload), &nodeMsg)
			if err != nil {
				logger.Errorw("error at parsing redis message", err)
				continue
			}
			m.msgChan <- &nodeMsg
		}
	}
}
