package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"load_balancer/backend"
	"load_balancer/backend/server_pool"
	"load_balancer/frontend"
	"load_balancer/utils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := utils.NewLogger()
	defer logger.Sync()

	config, err := utils.GetLBConfig()
	if err != nil {
		utils.Logger.Fatal(err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// now with config, create server pool
	serverPool, err := server_pool.NewServerPool(utils.GetLBStrategy(config.LBAlgorithm))
	if err != nil {
		utils.Logger.Fatal(err.Error())
	}

	loadBalancer := frontend.NewLoadBalancer(serverPool)

	for _, b := range config.BackendServers {
		u, err := url.Parse(b)
		if err != nil {
			logger.Fatal(err.Error(), zap.String("URL", b))
		}

		rp := httputil.NewSingleHostReverseProxy(u)
		backendServer := backend.NewBackend(u, rp)
		rp.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
			logger.Error("error handling the request",
				zap.String("host", u.Host),
				zap.Error(e),
			)
			backendServer.SetActive(false)

			if !frontend.AllowRetry(request) {
				utils.Logger.Info(
					"Max retry attempts reached, terminating",
					zap.String("address", request.RemoteAddr),
					zap.String("path", request.URL.Path),
				)
				http.Error(writer, "Service not available", http.StatusServiceUnavailable)
				return
			}

			logger.Info(
				"Attempting retry",
				zap.String("address", request.RemoteAddr),
				zap.String("URL", request.URL.Path),
				zap.Bool("retry", true),
			)

			loadBalancer.Serve(
				writer,
				request.WithContext(
					context.WithValue(request.Context(), frontend.RETRY_ATTEMPTED, true),
				),
			)
		}

		serverPool.AddBackend(backendServer)
	}

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", config.LBPort),
		Handler:           http.HandlerFunc(loadBalancer.Serve),
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("ListenAndServe() error", zap.Error(err))
	}
	logger.Info(
		"Load Balancer started", zap.Int("port", config.LBPort),
	)

	go server_pool.StartHealthCheck(ctx, serverPool)
	go func() {
		<-ctx.Done()
		shutdownCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}
	}()
}
