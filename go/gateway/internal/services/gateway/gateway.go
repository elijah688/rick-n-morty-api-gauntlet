package gateway

import (
	"net/http"
	"riki_gateway/internal/config"
)

type GatewayService struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *GatewayService {

	return &GatewayService{cfg, new(http.Client)}
}
