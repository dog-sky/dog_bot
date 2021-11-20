package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppPort string `envconfig:"APP_PORT" required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("DOG", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
