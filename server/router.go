package server

import (
	"net/http"
	"time"

	"github.com/anodamobi/go-tg-api/server/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/sirupsen/logrus"
)

const durationThreshold = time.Second * 10

func Router(
	log *logrus.Entry,
) chi.Router {

	router := chi.NewRouter()

	//TODO: update CORS policy
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposedHeaders:   []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		cors.Handler,
		middleware.Recoverer,
		middlewares.Logger(log, durationThreshold),
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Glance API"))
	})

	return router
}
