package server_pool

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"load_balancer/backend"
	"load_balancer/utils"
	"time"
)

// IServerPool defines all methods of a group of backend servers, which are untouched by properties of LB.
type IServerPool interface {
	GetServerPoolSize() int
	GetBackendServers() []backend.Backend
	GetNextValidPeer() backend.Backend
	AddBackend(backend.Backend)
}

func NewServerPool(strategy utils.LoadBalancer) (IServerPool, error) {
	switch strategy {
	case utils.RoundRobin:
		{
			return &roundRobinServerPool{}, nil
		}
	case utils.LeastConnections:
		{
			return &leastConnectionsServerPool{}, nil
		}
	default:
		return nil, fmt.Errorf("invalid strategy")
	}
}

// StartHealthCheck method is a worker which runs indefinitely
func StartHealthCheck(ctx context.Context, sp IServerPool) {
	t := time.NewTimer(time.Second * 20)
	utils.Logger.Info("Starting Health Check ...")

	for {
		select {
		case <-t.C:
			{
				go HealthCheck(ctx, sp) // by the time health check is called, server pool (strategy is finalised and created)
			}
		case <-ctx.Done(): // cancellation signal to close HealthCheck, so we need to pass context
			{
				utils.Logger.Info("Closing HealthCheck ...")
				return
			}
		}
	}
}

func HealthCheck(ctx context.Context, sp IServerPool) {
	aliveChanel := make(chan bool, 1)

	for _, b := range sp.GetBackendServers() {
		b := b
		requestContext, stop := context.WithTimeout(ctx, 10*time.Second) // request waits for 10 second for response to come back
		defer stop()

		go backend.IsBackendAlive(requestContext, aliveChanel, b.GetUrl())
		status := "up"

		select {
		case <-ctx.Done():
			{
				utils.Logger.Info("Gracefully shutting down health check")
				return
			}
		case alive := <-aliveChanel:
			{
				b.SetActive()
				if !alive {
					status = "down"
				}
			}
		}

		utils.Logger.Debug(
			"URL Status",
			zap.String("URL", b.GetUrl().String()),
			zap.String("status", status))
	}
}
