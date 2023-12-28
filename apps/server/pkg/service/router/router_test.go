package router_test

import (
	"context"
	"testing"
	"time"

	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/router/routerfakes"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/service/storage/storagefakes"
	"github.com/stretchr/testify/assert"
)

func TestRouterMustNotStartTwice(t *testing.T) {
	r := createTestRouter()

	_, err := r.Start()
	assert.NoError(t, err)

	_, err = r.Start()
	assert.Error(t, err)
}

func TestRouterMustNotStopIfIsNotRunning(t *testing.T) {
	r := createTestRouter()

	err := r.Stop()
	assert.Error(t, err)

	_, err = r.Start()
	assert.NoError(t, err)
	err = r.Stop()
	assert.NoError(t, err)
}

func TestStatsWorkerMustUpdate(t *testing.T) {
	// overwrite
	router.StatsTickerInterval = time.Second * 1

	r := createTestRouter()
	node, err := r.Start()
	assert.NoError(t, err)

	updatedAt1 := node.Stats.UpdatedAt

	time.Sleep(time.Second * 1)
	time.Sleep(time.Millisecond * 100)

	updatedAt2 := r.GetLocalNode().Stats.UpdatedAt

	assert.NotEqual(t, updatedAt1, updatedAt2)
}

func TestRouterMustDeleteNodeAfterStop(t *testing.T) {
	// overwrite
	router.StatsTickerInterval = time.Second * 1
	conf := &config.Config{Router: &config.RouterConfig{Region: "region-test"}}
	store := storage.NewLocalStorage()
	mon := &routerfakes.FakeMonitor{}
	r := router.NewRouter(conf, store, mon)

	_, err := r.Start()
	assert.NoError(t, err)

	nodes, err := store.ListNodes(context.Background(), "region-test")
	t.Log(nodes)
	assert.Equal(t, 1, len(nodes))

	time.Sleep(time.Second * 1)
	err = r.Stop()

	nodes, err = store.ListNodes(context.Background(), "region-test")
	t.Log(nodes)
	assert.Equal(t, 0, len(nodes))
}

func createTestRouter() *router.Router {
	conf := &config.Config{Router: &config.RouterConfig{}}
	store := &storagefakes.FakeInMemoryStorage{}
	mon := &routerfakes.FakeMonitor{}

	return router.NewRouter(conf, store, mon)
}
