package frontend

import (
	"load_balancer/backend/server_pool"
	"net/http"
)

// ILoadBalancer interface - hi use krege FE ke log, think of this as UI, and all configs you've entered in config.yaml
type ILoadBalancer interface {
	Serve(http.ResponseWriter, *http.Request)
}

type loadBalancer struct {
	serverPool server_pool.IServerPool
}

func (l *loadBalancer) Serve(w http.ResponseWriter, r *http.Request) {
	peer := l.serverPool.GetNextValidPeer()
	if peer != nil {
		peer.Serve(w, r)
		return
	}

	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func NewLoadBalancer(serverPool server_pool.IServerPool) ILoadBalancer {
	return &loadBalancer{serverPool: serverPool}
}
