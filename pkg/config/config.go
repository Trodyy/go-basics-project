package config

import (
	"github.com/caarlos0/env/v6"
)

type Server struct {
	Host string `env:"SERVERHOST"`
	Port int `env:"SERVERPORT"`
}


type Postgres struct {
	User   string `env:"PGUSER"`
	Pass   string `env:"PGPASS"`
	Host   string `env:"PGHOST"`
	Port   int    `env:"PGPORT"`
	DbName string `env:"PGDB"`
}


type Config struct {
	Server
	Postgres
}


func LoadConfigOrPanic() Config {
	var config *Config = new(Config)
	if err := env.Parse(config); err != nil {
		panic(err)
	}

	return *config
}