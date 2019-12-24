package main

import (
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

	cr := payment.NewCieloRepository(
		payment.NewHTTPRequester(
			logger,
			config.CieloURI,
			map[string]string{
				"merchantid":  config.CieloMerchantID,
				"merchantkey": config.CieloMerchantKey,
			},
			config.GeneralReqTimeout,
		),
	)

	cs := payment.NewCieloStrategy(cr)
	re := payment.NewRedeStrategy()

	a := payment.NewAcquirerProvider(
		payment.AcquirersStrategy{
			payment.Cielo: cs,
			payment.Rede:  re,
		},
	)
	r := payment.NewSourcesRepository(&payment.LoggableDBWrapper{DB: db, Logger: logger})
	s := payment.NewService(r, a)

	handler := createServerHandler(s)

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
