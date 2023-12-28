package router

import (
	"sync"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type monitor struct {
	mu      sync.RWMutex
	cpuPrev *cpu.Stats
}

func NewMonitor() (Monitor, error) {
	stats, err := cpu.Get()

	if err != nil {
		return nil, err
	}

	return &monitor{
		cpuPrev: stats,
		mu:      sync.RWMutex{},
	}, nil
}

func (m *monitor) GetCpuUsage() (float32, error) {
	next, err := cpu.Get()

	if err != nil {
		return 0, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	totalDiff := float32(next.Total - m.cpuPrev.Total)
	idleDiff := float32(next.Idle - m.cpuPrev.Idle)

	m.cpuPrev = next

	if totalDiff < 0 {
		return 0, nil
	}

	cpuUsage := (totalDiff - idleDiff) / totalDiff * 100.0
	return cpuUsage, nil
}

func (m *monitor) GetMemUsage() (float32, error) {
	stats, err := memory.Get()

	if err != nil {
		return 0, err
	}

	memUsage := (float32(stats.Used) / float32(stats.Total)) * 100.0

	return memUsage, nil
}
