package config

type Config struct {
	GatewayConfig *GatewayConfig
}

func New() (*Config, error) {

	db, err := newGatewayConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		GatewayConfig: db,
	}, nil

}
