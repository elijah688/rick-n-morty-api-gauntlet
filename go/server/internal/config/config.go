package config

type Config struct {
	RikiClientConfig *RikiClientConfig
	DBConfig         *DBConfig
}

func New() (*Config, error) {

	db, err := newDBConfig()
	if err != nil {
		return nil, err
	}

	riki, err := newRikiClientConfig()
	if err != nil {
		return nil, err
	}
	return &Config{
		RikiClientConfig: riki,
		DBConfig:         db,
	}, nil

}
