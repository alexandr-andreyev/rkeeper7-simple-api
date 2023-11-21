package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/sys/windows/svc/debug"
)

type Server struct {
	Winlog debug.Log
	// a local logger
	Logger *slog.Logger
	// a database connection

	// your app configuration
	Config *Config `yaml:"config"`
}

type Config struct {
	// текущее окружение local, dev, prod
	Env             string           `yaml:"env" env-default:"local"`
	ServerPort      int              `yaml:"server_port" env-default:"9081"`
	RK7ClientConfig *RK7ClientConfig `yaml:"rk7_client_config"`
}

type RK7ClientConfig struct {
	CashServerIp   string `yaml:"cash_server_ip"`
	CashServerPort int    `yaml:"cash_server_port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	//Timeout        time.Duration `yaml:"timeout"`
}

var ServerConfig Server

func New() (Config, error) {
	configPath := "./configs/config.yml"
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	fmt.Printf("Config: %+v\n", cfg)
	return cfg, nil
}
