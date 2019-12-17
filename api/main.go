package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"payments-processor/payment-processor"

	"database/sql"

	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/go-sql-driver/mysql"
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

	db, oErr := sql.Open(config.DBDriver, config.DBConnectionString)
	if oErr != nil {
		log.Fatal(lErr)
	}

	r := payment.NewSourcesRepository(db)

	fmt.Println(r.GetByID(context.Background(), "token-1"))

	handler := createServerHandler()

	logger.Info(
		fmt.Sprintf("starting application at port: %d", config.AppPort),
	)

	gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppPort),
		Handler:      handler,
		ReadTimeout:  config.AppReadTimeout,
		WriteTimeout: config.AppWriteTimeout,
	})
}
