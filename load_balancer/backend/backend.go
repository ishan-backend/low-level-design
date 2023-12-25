package backend

import (
	"context"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend interface {
	GetIsAlive() bool
	SetActive(bool)
	GetActiveConnections() int
	GetUrl() *url.URL
	Serve(http.ResponseWriter, *http.Request)
}

func IsBackendAlive(ctx context.Context, aliveChannel chan bool, u *url.URL) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", u.Host)
	if err != nil {
		utils.Logger.Debug("Site unreachable", zap.Error(err))
		aliveChannel <- false
		return
	}
	_ = conn.Close()
	aliveChannel <- true
}

type backend struct {
	url          *url.URL // ip
	alive        bool
	connections  int
	reverseProxy *httputil.ReverseProxy
	mux          sync.RWMutex
}

func (b *backend) GetIsAlive() bool {
	b.mux.Lock()
	isAlive := b.alive
	b.mux.Unlock()
	return isAlive
}

func (b *backend) SetActive(val bool) {
	b.mux.Lock()
	b.alive = val
	b.mux.Unlock()
}

func (b *backend) GetActiveConnections() int {
	b.mux.RLock()
	connections := b.connections
	b.mux.RUnlock()
	return connections
}

func (b *backend) GetUrl() *url.URL {
	return b.url
}

func (b *backend) Serve(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		b.mux.Lock()
		b.connections--
		b.mux.Unlock()
	}()

	b.mux.Lock()
	b.connections++
	b.mux.Unlock()
	// now connections updated value will be reflected till serve http is executed
	b.reverseProxy.ServeHTTP(writer, request)
}

// NewBackend is called to add a new backend server to server pool
func NewBackend(u *url.URL, reverseProxy *httputil.ReverseProxy) Backend {
	return &backend{
		url:          u,
		alive:        true,
		reverseProxy: reverseProxy,
	}
}
