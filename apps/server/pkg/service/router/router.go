package router

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/storage"
)

var (
	StatsTickerInterval = time.Second * 2
)

type Router struct {
	mu        sync.RWMutex
	ctx       context.Context
	conf      *config.RouterConfig
	store     storage.ObjectStore
	doneChan  chan struct{}
	localNode *core.Node
	running   atomic.Bool
	monitor   Monitor
	messenger Messenger

	onNodeMessage func(msg *core.NodeMessage) error
}

func NewRouter(conf *config.Config, localNode *core.Node, store storage.ObjectStore, monitor Monitor, messenger Messenger) *Router {
	return &Router{
		mu:        sync.RWMutex{},
		ctx:       context.Background(),
		store:     store,
		conf:      conf.Router,
		doneChan:  make(chan struct{}),
		running:   atomic.Bool{},
		monitor:   monitor,
		messenger: messenger,
		localNode: localNode,
	}
}

func (r *Router) OnNodeMessage(f func(msg *core.NodeMessage) error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.onNodeMessage = f
}

func (r *Router) Start() (*core.Node, error) {
	if r.running.Load() {
		return nil, errors.New("already running")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// store local node
	if err := r.store.StoreNode(r.ctx, r.localNode.Region, r.localNode); err != nil {
		return nil, err
	}

	go r.statsWorker()
	go r.messageWorker()

	r.running.Store(true)

	return r.localNode, nil
}

func (r *Router) Stop() error {
	if !r.running.Load() {
		return errors.New("router is not running")
	}

	close(r.doneChan)

	r.mu.RLock()
	node := r.localNode
	r.mu.RUnlock()

	if err := r.store.DeleteNode(r.ctx, node.Region, node.Id); err != nil {
		return err
	}

	return nil
}

func (r *Router) GetLocalNode() *core.Node {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.localNode
}

func (r *Router) updateNodeStats() error {
	cpuUsage, err := r.monitor.GetCpuUsage()
	memUsage, err := r.monitor.GetMemUsage()

	if err != nil {
		return err
	}

	r.mu.Lock()
	r.localNode.Stats.UpdatedAt = time.Now().Unix()
	r.localNode.Stats.CpuLoad = cpuUsage
	r.localNode.Stats.MemoryLoad = memUsage
	r.mu.Unlock()
	return nil
}

func (r *Router) GetAvaliableNode() (*core.Node, error) {
	nodes, err := r.store.ListNodes(r.ctx, r.localNode.Region)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, NotFoundAvaliableNode
	}

	var avaliable *core.Node

	for _, n := range nodes {
		// filter nodes with greater cpu 80% of cpu usage and 80% o memory usage
		if n.Stats.CpuLoad > 80 || n.Stats.MemoryLoad > 90 {
			continue
		}

		avaliable = n
	}

	if avaliable == nil {
		return nil, NotFoundAvaliableNode
	}

	return avaliable, nil
}

func (r *Router) SendNodeMessage(nodeID string, msg *core.NodeMessage) error {
	return r.messenger.WriteMessage(nodeID, msg)
}

func (r *Router) statsWorker() {
	ticker := time.NewTicker(StatsTickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-r.doneChan:
			return
		case <-ticker.C:
			if !r.running.Load() {
				return
			}

			r.updateNodeStats()

			r.mu.RLock()
			err := r.store.StoreNode(r.ctx, r.localNode.Region, r.localNode)

			if err != nil {
				logger.Errorw("error at storing new stats for local node", err)
			}
			r.mu.RUnlock()
		}
	}
}

func (r *Router) messageWorker() {
	for {
		select {
		case <-r.doneChan:
			return
		case msg := <-r.messenger.ReadChan():
			if r.onNodeMessage == nil {
				continue
			}

			if err := r.onNodeMessage(msg); err != nil {
				logger.Errorw("an error ocurred at handling node message", err)
			}
		}
	}
}
