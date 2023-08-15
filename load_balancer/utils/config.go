package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type LoadBalancer int

const (
	RoundRobin LoadBalancer = iota
	LeastConnections
)

type Config struct {
	LBPort                     int      `yaml:"lb_port"`
	LBAlgorithm                string   `yaml:"lb_algorithm"`
	BackendServers             []string `yaml:"backend_servers"`
	MaxAttemptsCheckAliveLimit int      `yaml:"max_attempts_check_alive_limit"`
}

const MAX_ATTEMPTS_CHECK_ALIVE_LIMIT int = 5

func GetLBConfig() (*Config, error) {
	config := Config{}
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	// Validations
	if len(config.BackendServers) == 0 {
		return nil, errors.New("backend hosts expected, none provided")
	}

	if config.LBPort == 0 {
		return nil, errors.New("load balancer port not found")
	}

	// Responsibility of client to provide lb strategy
	return &config, nil
}

// GetLBStrategy function translates config value to enumerated type LoadBalancer to pass to NewServerPool
func GetLBStrategy(strategy string) LoadBalancer {
	switch strategy {
	case "least-connection":
		return LeastConnections
	default:
		return RoundRobin
	}
}
