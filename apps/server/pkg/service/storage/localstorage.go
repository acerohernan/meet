package storage

import (
	"context"
	"sync"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

type LocalStorage struct {
	mu sync.RWMutex
	// map of region -> map of nodeID -> Node
	nodes map[string]map[string]*core.Node
}

func NewLocalStorage() *LocalStorage {
	logger.Infow("using local storage as in-memory storage")
	return &LocalStorage{
		mu:    sync.RWMutex{},
		nodes: make(map[string]map[string]*core.Node),
	}
}

func (s *LocalStorage) StoreNode(ctx context.Context, region string, node *core.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.nodes[region] == nil {
		s.nodes[region] = make(map[string]*core.Node)
	}

	s.nodes[region][node.Id] = node

	return nil
}

func (s *LocalStorage) ListNodes(ctx context.Context, region string) ([]*core.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodes := s.nodes[region]

	if nodes == nil {
		return nil, nil
	}

	items := make([]*core.Node, 0)

	for _, n := range nodes {
		items = append(items, n)
	}

	return items, nil
}

func (s *LocalStorage) DeleteNode(ctx context.Context, region string, nodeID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.nodes[region] == nil {
		return nil
	}

	delete(s.nodes[region], nodeID)

	return nil
}
