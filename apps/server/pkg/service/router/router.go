package router

import (
	"context"
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
)

var (
	StatsTickerInterval = time.Second * 2
)

type Router struct {
	mu        sync.RWMutex
	ctx       context.Context
	conf      *config.RouterConfig
	store     storage.InMemoryStorage
	doneChan  chan struct{}
	localNode *core.Node
	running   atomic.Bool
	monitor   Monitor
}

func NewRouter(conf *config.Config, store storage.InMemoryStorage, monitor Monitor) *Router {
	return &Router{
		mu:       sync.RWMutex{},
		ctx:      context.Background(),
		store:    store,
		conf:     conf.Router,
		doneChan: make(chan struct{}),
		running:  atomic.Bool{},
		monitor:  monitor,
	}
}

func (r *Router) Start() (*core.Node, error) {
	if r.running.Load() {
		return nil, errors.New("already running")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.localNode = createLocalNode(r.conf.Region)

	// store local node
	if err := r.store.StoreNode(r.ctx, r.localNode.Region, r.localNode); err != nil {
		return nil, err
	}

	go r.statsWorker()

	r.running.Store(true)

	return r.localNode, nil
}

func (r *Router) Stop() error {
	if !r.running.Load() {
		return errors.New("router is not running")
	}

	close(r.doneChan)

	r.mu.RLock()
	if err := r.store.DeleteNode(r.ctx, r.localNode.Region, r.localNode.Id); err != nil {
		return err
	}
	r.mu.RUnlock()

	return nil
}

func (r *Router) GetLocalNode() *core.Node {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.localNode
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
				logger.Errow("error at storing new stats for local node", err)
			}
			r.mu.RUnlock()
		}
	}
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

func createLocalNode(region string) *core.Node {
	return &core.Node{
		Id:     utils.NewId(utils.NodePrefix),
		Region: region,
		Stats: &core.NodeStats{
			StartedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			NumCpus:   uint32(runtime.NumCPU()),
		},
	}
}
