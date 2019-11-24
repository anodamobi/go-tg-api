package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/anodamobi/go-tg-api/bot"

	"github.com/pkg/errors"

	"github.com/anodamobi/go-tg-api/config"
	"github.com/anodamobi/go-tg-api/server"

	"github.com/sirupsen/logrus"
)

type App struct {
	config config.Config
	log    *logrus.Entry
}

func New(config config.Config) *App {
	return &App{
		config: config,
		log:    config.Log(),
	}
}

func (a *App) Start() error {
	conf := a.config
	log := conf.Log()

	botBoss, err := bot.NewBoss(
		conf.Bot(),
		conf.DB(),
		log,
		conf.JWT(),
	)
	if err != nil {
		return errors.Wrap(err, "failed to connect with bot")
	}

	go func() {
		log.Info("bot starting")
		err := botBoss.Listen()
		if err != nil {
			panic(err)
		}
	}()

	httpConfiguration := conf.HTTP()
	router := server.Router(
		log,
		botBoss.Summary(),
	)

	serverHost := fmt.Sprintf("%s:%s", httpConfiguration.Host, httpConfiguration.Port)
	a.log.WithField("api", "start").
		Info(fmt.Sprintf("listenig addr =  %s, tls = %v", serverHost, httpConfiguration.SSL))

	httpServer := http.Server{
		Addr:           serverHost,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	switch httpConfiguration.SSL {
	case true:
		tlsConfig := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
			},
			InsecureSkipVerify: true,
		}

		httpServer.TLSConfig = tlsConfig
		if err := httpServer.ListenAndServeTLS(httpConfiguration.ServerCertPath, httpConfiguration.ServerKeyPath); err != nil {
			return errors.Wrap(err, "failed to start https server")
		}

	default:
		if err := httpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "failed to start http server")
		}
	}

	return nil
}
