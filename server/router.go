package server

import (
	"time"

	"github.com/anodamobi/go-tg-api/server/handlers"

	"github.com/anodamobi/go-tg-api/server/middlewares"

	"github.com/anodamobi/go-tg-api/bot"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/sirupsen/logrus"
)

const durationThreshold = time.Second * 10

func Router(
	log *logrus.Entry,
	botSummary bot.Summary,
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
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Recoverer,
		middlewares.Logger(log, durationThreshold),
	)

	router.Get("/bot", handlers.NewHomeHandler(botSummary, log).Handle)

	return router
}
