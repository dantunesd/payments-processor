package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"payments-processor/payment-processor"

	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	cfg, cErr := NewConfig()
	if cErr != nil {
		log.Fatal(cErr)
	}

	l, lErr := zap.NewProduction()
	if lErr != nil {
		log.Fatal(lErr)
	}
	defer l.Sync()

	cr := payment.NewCieloRepository(
		payment.NewHTTPRequester(
			&http.Client{Timeout: cfg.GeneralReqTimeout},
			l,
			cfg.CieloURI,
			map[string]string{
				"merchantid":  cfg.CieloMerchantID,
				"merchantkey": cfg.CieloMerchantKey,
			},
		),
	)

	cs := payment.NewCieloStrategy(cr)

	rr := payment.NewRedeRepository(
		payment.NewHTTPRequester(
			&http.Client{Timeout: cfg.GeneralReqTimeout},
			l,
			cfg.RedeURI,
			map[string]string{
				"Authorization": cfg.RedeAuth,
			},
		),
	)

	rs := payment.NewRedeStrategy(rr)

	ap := payment.NewAcquirerProvider(
		payment.AcquirerStrategies{
			payment.Cielo: cs,
			payment.Rede:  rs,
		},
	)

	db, oErr := sql.Open(cfg.DBDriver, cfg.DBConnectionString)
	if oErr != nil {
		l.Fatal(lErr.Error())
	}

	sr := payment.NewSourcesRepository(
		payment.NewLoggableDBWrapper(db, l),
	)

	s := payment.NewService(sr, ap)

	l.Info(fmt.Sprintf("starting application at port: %d", cfg.AppPort))

	gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.AppPort),
		Handler:      createServerHandler(s),
		ReadTimeout:  cfg.AppReadTimeout,
		WriteTimeout: cfg.AppWriteTimeout,
	})
}
