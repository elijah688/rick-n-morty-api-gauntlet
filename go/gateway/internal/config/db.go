// config/config.go
package config

import (
	"fmt"
	"os"
)

type GatewayConfig struct {
	CRUDServiceHost   string
	GatewayServerPort string
}

func newGatewayConfig() (*GatewayConfig, error) {
	crucSvcHost, err := getEnv("CRUD_SVC_HOST")
	if err != nil {
		return nil, err
	}
	gatewayServerPort, err := getEnv("GATEWAY_SERVER_PORT")
	if err != nil {
		return nil, err
	}

	return &GatewayConfig{
		CRUDServiceHost:   crucSvcHost,
		GatewayServerPort: gatewayServerPort,
	}, nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
