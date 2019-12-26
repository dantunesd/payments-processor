package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is a representation for env variables
type Config struct {
	AppName            string        `envconfig:"APP_NAME" default:"Payments-Processor"`
	AppPort            int           `envconfig:"APP_PORT" default:"3000"`
	AppReadTimeout     time.Duration `envconfig:"APP_READ_TIMEOUT" default:"6s"`
	AppWriteTimeout    time.Duration `envconfig:"APP_WRITE_TIMEOUT" default:"6s"`
	DBDriver           string        `envconfig:"DB_DRIVER" default:"mysql"`
	DBConnectionString string        `envconfig:"DB_CONNECTION_STRING" default:"root:root@/source"`
	CieloURI           string        `envconfig:"CIELO_URI" default:"http://localhost:8010"`
	CieloMerchantID    string        `envconfig:"CIELO_MERCHANT_ID" default:"f85ca5cc-335a-4dff-9ed1-4d500cd21bbd"`
	CieloMerchantKey   string        `envconfig:"CIELO_MERCHANT_KEY" default:"AUZGAZLATBVIEMEFFCJVWVDPGWZBSXDYREUESDYJ"`
	RedeURI            string        `envconfig:"REDE_URI" default:"http://localhost:8010"`
	RedeAuth           string        `envconfig:"REDE_AUTH" default:"Basic MTAwMDQ5NzI6NzM2MzRhNTE3NzE4NGY0NDk3NTMwYTU0NGZlMmZiOWM="`
	GeneralReqTimeout  time.Duration `envconfig:"GENERAL_REQ_TIMEOUT" default:"4s"`
	// RedeURI            string        `envconfig:"REDE_URI" default:"https://api.userede.com.br/desenvolvedores"`
	// CieloURI           string        `envconfig:"CIELO_URI" default:"https://apisandbox.cieloecommerce.cielo.com.br"`
}

// NewConfig Config constructor.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, envconfig.Process("", cfg)
}
