package router

import (
	"runtime"
	"time"

	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/utils"
)

func CreateLocalNode(conf *config.RouterConfig) *core.Node {
	return &core.Node{
		Id:     utils.NewId(utils.NodePrefix),
		Region: conf.Region,
		Stats: &core.NodeStats{
			StartedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			NumCpus:   uint32(runtime.NumCPU()),
		},
	}
}
