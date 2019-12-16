package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"go.uber.org/zap"
)

func main() {
	config, cErr := NewConfig()
	if cErr != nil {
		log.Fatal(cErr)
	}

	logger, lErr := zap.NewProduction()
	if lErr != nil {
		log.Fatal(lErr)
	}
	defer logger.Sync()

	handler := createServerHandler()

	gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppPort),
		Handler:      handler,
		ReadTimeout:  config.AppReadTimeout,
		WriteTimeout: config.AppWriteTimeout,
	})
}
