package config

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sys/windows/svc/debug"
)

type Server struct {
	Winlog debug.Log
	// a local logger
	// a database connection
	// your app configuration
	Config *Config
}

type Config struct {
	ServerPort            int
	RKeeperCashServerIp   string
	RkeeperCashServerPort int
	RkeeperUser           string
	RkeeperPassword       string
}

var ServerConfig Server

func New() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("api", cfg); err != nil {
		fmt.Println("envconfig process:", err)
		return nil, err
	}
	ServerConfig.Config = cfg
	return cfg, nil
}
