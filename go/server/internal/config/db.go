// config/config.go
package config

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	MainDBHost string
	MainDBPort string
	MainDBUser string
	MainDBPass string
	MainDBName string

	AppDBName string
}

func newDBConfig() (*DBConfig, error) {
	mainDBHost, err := getEnv("MAIN_DB_HOST")
	if err != nil {
		return nil, err
	}
	mainDBPort, err := getEnv("MAIN_DB_PORT")
	if err != nil {
		return nil, err
	}
	mainDBUser, err := getEnv("MAIN_DB_USER")
	if err != nil {
		return nil, err
	}
	mainDBPass, err := getEnv("MAIN_DB_PASS")
	if err != nil {
		return nil, err
	}
	mainDBName, err := getEnv("MAIN_DB_NAME")
	if err != nil {
		return nil, err
	}

	appDBName, err := getEnv("APP_DB_NAME")
	if err != nil {
		return nil, err
	}

	return &DBConfig{
		MainDBHost: mainDBHost,
		MainDBPort: mainDBPort,
		MainDBUser: mainDBUser,
		MainDBPass: mainDBPass,
		MainDBName: mainDBName,

		AppDBName: appDBName,
	}, nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}

func (c *DBConfig) MainDBConfig() (*pgxpool.Config, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.MainDBUser, c.MainDBPass, c.MainDBHost, c.MainDBPort, c.MainDBName)

	return pgxpool.ParseConfig(connString)
}

func (c *DBConfig) AppDBConfig() (*pgxpool.Config, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.MainDBUser, c.MainDBPass, c.MainDBHost, c.MainDBPort, c.AppDBName)

	return pgxpool.ParseConfig(connString)
}
