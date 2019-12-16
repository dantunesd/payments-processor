package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is a representation for env variables
type Config struct {
	AppName         string        `envconfig:"APP_NAME" default:"Payments-Processor"`
	AppPort         int           `envconfig:"APP_PORT" default:"3000"`
	AppReadTimeout  time.Duration `envconfig:"APP_READ_TIMEOUT" default:"3s"`
	AppWriteTimeout time.Duration `envconfig:"APP_WRITE_TIMEOUT" default:"3s"`
}

// NewConfig Config constructor.
func NewConfig() (*Config, error) {
	config := &Config{}
	return config, envconfig.Process("", config)
}
