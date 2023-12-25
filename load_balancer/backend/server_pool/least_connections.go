package server_pool

import (
	"load_balancer/backend"
	"sync"
)

type leastConnectionsServerPool struct {
	backend []backend.Backend
	mux     sync.RWMutex
}

func (l *leastConnectionsServerPool) GetServerPoolSize() int {
	return len(l.backend)
}

func (l *leastConnectionsServerPool) GetBackendServers() []backend.Backend {
	return l.backend
}

// todo: find out how can this fail?
func (l *leastConnectionsServerPool) GetNextValidPeer() backend.Backend {
	var leastConnectedPeer backend.Backend
	for i := 0; i < len(l.backend); i++ {
		if l.backend[i].GetIsAlive() {
			leastConnectedPeer = l.backend[i]
			break
		}
	}
	for _, b := range l.backend {
		if !b.GetIsAlive() {
			continue
		}
		if b.GetActiveConnections() < leastConnectedPeer.GetActiveConnections() {
			leastConnectedPeer = b
		}
	}

	return leastConnectedPeer
}

func (l *leastConnectionsServerPool) AddBackend(backend backend.Backend) {
	l.backend = append(l.backend, backend)
}
