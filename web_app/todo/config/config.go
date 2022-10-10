package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
