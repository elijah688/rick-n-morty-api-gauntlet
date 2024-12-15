package services

import (
	"riki_gateway/internal/config"
	"riki_gateway/internal/services/gateway"
)

type Services struct {
	gateway *gateway.GatewayService
	cfg     *config.Config
}

func New(cfg *config.Config) *Services {
	return &Services{
		cfg:     cfg,
		gateway: gateway.New(cfg)}
}

func (svcs *Services) Gateway() *gateway.GatewayService {
	return svcs.gateway
}

func (svcs *Services) Config() *config.Config {
	return svcs.cfg
}
