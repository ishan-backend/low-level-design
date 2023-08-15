package server_pool

import (
	"load_balancer/backend"
	"sync"
)

type IRoundRobin interface {
	IServerPool
	Rotate() backend.Backend
}

type roundRobinServerPool struct {
	backend []backend.Backend
	mux     sync.RWMutex
	current int
}

func (r *roundRobinServerPool) GetServerPoolSize() int {
	return len(r.backend)
}

func (r *roundRobinServerPool) GetBackendServers() []backend.Backend {
	return r.backend
}

// GetNextValidPeer checks if a backend server from pool is alive, and returns it to LB if it is alive
func (r *roundRobinServerPool) GetNextValidPeer() backend.Backend {
	// run n times to check next available peer
	for i := 0; i < r.GetServerPoolSize(); i++ {
		nextPeer := r.Rotate() // update index to check for next peer
		if nextPeer.GetIsAlive() {
			return nextPeer
		}
	}

	return nil
}

func (r *roundRobinServerPool) AddBackend(backend backend.Backend) {
	r.backend = append(r.backend, backend)
}

// Rotate rotates current to next backend server, dumb way of RR algo
func (r *roundRobinServerPool) Rotate() backend.Backend {
	// lock is required where critical section can be accessed concurrently and may be modified
	// two requests might come to LB, then current for roundRobin has to update, in such case CS will be updated only via exclusive lock.
	r.mux.Lock()
	r.current = (r.current + 1) % r.GetServerPoolSize()
	r.mux.Unlock()
	return r.backend[r.current]
}
