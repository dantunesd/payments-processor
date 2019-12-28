package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config holds the environment variables values.
type Config struct {
	AppName            string        `envconfig:"APP_NAME" default:"Payments-Processor"`
	AppPort            int           `envconfig:"APP_PORT" default:"3000"`
	AppReadTimeout     time.Duration `envconfig:"APP_READ_TIMEOUT" default:"10s"`
	AppWriteTimeout    time.Duration `envconfig:"APP_WRITE_TIMEOUT" default:"10s"`
	DBDriver           string        `envconfig:"DB_DRIVER" default:"mysql"`
	DBConnectionString string        `envconfig:"DB_CONNECTION_STRING" default:"root:root@/source"`
	CieloURI           string        `envconfig:"CIELO_URI" default:"http://localhost:8010"`
	CieloMerchantID    string        `envconfig:"CIELO_MERCHANT_ID" default:"merchant-id"`
	CieloMerchantKey   string        `envconfig:"CIELO_MERCHANT_KEY" default:"merchant-key"`
	RedeURI            string        `envconfig:"REDE_URI" default:"http://localhost:8010"`
	RedeAuth           string        `envconfig:"REDE_AUTH" default:"Basic authentication"`
	GeneralReqTimeout  time.Duration `envconfig:"GENERAL_REQ_TIMEOUT" default:"6s"`
}

// NewConfig Config's constructor.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, envconfig.Process("", cfg)
}
