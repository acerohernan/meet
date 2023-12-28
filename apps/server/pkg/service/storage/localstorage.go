package storage

import (
	"context"
	"sync"

	"github.com/acerohernan/meet/core"
)

type LocalStorage struct {
	mu sync.RWMutex
	// map of nodeID -> Node
	nodes map[string]*core.Node
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		mu:    sync.RWMutex{},
		nodes: make(map[string]*core.Node),
	}
}

func (s *LocalStorage) StoreNode(ctx context.Context, node *core.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nodes[node.Id] = node

	return nil
}

func (s *LocalStorage) ListNodes(ctx context.Context) ([]*core.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nodes := make([]*core.Node, 0)

	for _, n := range s.nodes {
		nodes = append(nodes, n)
	}

	return nodes, nil
}

func (s *LocalStorage) DeleteNode(ctx context.Context, nodeID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.nodes, nodeID)

	return nil
}
