package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
	DB     DBConfig
}

type DBConfig struct {
	Name     string `envconfig:"MYSQL_DATABASE" default:"code_kakitai"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASS" default:""`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Host     string `envconfig:"DB_HOST" default:"127.0.0.1"`
}

type Server struct {
	Address string `envconfig:"ADDRESS" default:""`
	Port    string `envconfig:"PORT" default:"8080"`
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(err)
		}
	})
	envconfig.Process("", &config)
	return &config
}
