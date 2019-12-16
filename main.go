package main

import (
	"go.uber.org/zap"
	"log"
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

	log.Println(config)
}
