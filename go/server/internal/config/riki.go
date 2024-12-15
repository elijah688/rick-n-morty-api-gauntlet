package config

import (
	"fmt"
	"os"
)

type RikiClientConfig struct {
	BaseURL string
}

func newRikiClientConfig() (*RikiClientConfig, error) {
	baseURL := os.Getenv("RIKI_API_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("RIKI_API_BASE_URL environment variable is not set")
	}

	return &RikiClientConfig{
		BaseURL: baseURL,
	}, nil
}
