package load_balancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// Backend is interface for interacting with a backend server
type Backend interface {
	IsAlive() bool
	SetAlive() bool

	GetActiveConnections() int
	GetUrl() *url.URL
	Serve(http.ResponseWriter, *http.Request) // relays request to corresponding url using reverse proxy
}

type backend struct {
	url          *url.URL
	alive        bool
	mux          sync.RWMutex
	connections  int
	reverseProxy *httputil.ReverseProxy
}

type ServerPool interface {
	GetServerPoolSize() int
	AddBackend(Backend)
	GetBackends() []Backend
	GetNextValidPeer() Backend
}

// Round Robin Implementation for server pool
type roundRobinServerPool struct {
	backendServers []Backend
	current        int
	mux            sync.RWMutex
}

func (r *roundRobinServerPool) Rotate() Backend {
	r.mux.Lock()
	r.current = (r.current + 1) %
		r.mux.Unlock()
	return r.backendServers[r.current]
}

func main() {

}
