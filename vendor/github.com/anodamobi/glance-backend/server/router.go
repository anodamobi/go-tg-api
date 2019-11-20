package server

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/anodamobi/glance-backend/server/handlers"
	"github.com/anodamobi/glance-backend/server/handlers/events"
	"github.com/anodamobi/glance-backend/server/handlers/maintenance"
	"github.com/anodamobi/glance-backend/server/handlers/notifications"
	"github.com/anodamobi/glance-backend/server/handlers/parcels"
	"github.com/anodamobi/glance-backend/server/handlers/user"
	"github.com/anodamobi/glance-backend/server/handlers/userevents"
	"github.com/anodamobi/glance-backend/server/handlers/wellbeing"
	"github.com/anodamobi/glance-backend/server/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/minio/minio-go"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const durationThreshold = time.Second * 10

func Router(
	log *logrus.Entry,
	db db.QInterface,
	auth *config.Authentication,
	oneSignal *config.OneSignal,
	s3 *minio.Client,
) chi.Router {

	notifications.NewEventNotificationWorker(db.DelayedNotificationsQ(), log, oneSignal).Check()

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
		middlewares.Ctx(
			handlers.CtxLog(log),
		),
	)

	router.Group(func(r chi.Router) {
		//r.Use(middlewares.Authenticator(auth.VerifyKey))

		r.Route("/user", func(r chi.Router) {
			r.Post("/", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).NewUserByAdmin)
			r.Get("/", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).GetUser)
			r.Patch("/", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).UpdateUser)
			r.Patch("/update_password", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).UpdatePassword)
			r.Post("/login", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).Login)
			r.Get("/temp_pass", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).GetTempPass)
			r.Post("/device", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).InsertDeviceID)
			r.Post("/send_code", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).SendCodeForRenewPassword)
			r.Post("/check_code", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).CheckForgetPwdCode)
			r.Post("/renew_pwd", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).RenewPassword)
			r.Get("/codes", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).GetCodes)

			r.Post("/admin", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).CreateAdmin)
			r.Patch("/admin", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).UpdateAdmin)
			r.Delete("/admin", user.New(db.UsersQ(), db.CodesForForgottenPwdQ(), db.AdminsQ(), log, auth).DeleteAdmin)
		})

		r.Route("/event", func(r chi.Router) {
			r.Post("/", events.New(db.EventsQ(), log, auth, db.UserEventsQ()).CreateEvent)
			r.Get("/", events.New(db.EventsQ(), log, auth, db.UserEventsQ()).GetEvent)
			r.Get("/all", events.New(db.EventsQ(), log, auth, db.UserEventsQ()).GetEvents)
			r.Patch("/filtered", events.New(db.EventsQ(), log, auth, db.UserEventsQ()).FilterEvents)
			r.Patch("/save", userevents.New(db.UsersQ(), log, auth, db.UserEventsQ(), oneSignal, db.DelayedNotificationsQ(), db.EventsQ()).SaveEvent)
			r.Patch("/unsave", userevents.New(db.UsersQ(), log, auth, db.UserEventsQ(), oneSignal, db.DelayedNotificationsQ(), db.EventsQ()).UnsaveEvent)
			r.Patch("/status", userevents.New(db.UsersQ(), log, auth, db.UserEventsQ(), oneSignal, db.DelayedNotificationsQ(), db.EventsQ()).ChangeStatus)
		})

		r.Route("/notification", func(r chi.Router) {
			r.Post("/", notifications.New(db.NotificationsQ(), db.UsersQ(), log, auth, oneSignal).CreateNotification)
			r.Get("/list", notifications.New(db.NotificationsQ(), db.UsersQ(), log, auth, oneSignal).ListNotifications)
		})

		r.Route("/wellbeing", func(r chi.Router) {
			r.Post("/", wellbeing.New(db.UsersQ(), log, auth, db.WellbeingRequestsQ()).CreateWellbeingRequest)
			r.Patch("/", wellbeing.New(db.UsersQ(), log, auth, db.WellbeingRequestsQ()).CloseWellbeingRequest)
		})

		r.Route("/parcel", func(r chi.Router) {
			r.Patch("/later", parcels.New(db.UsersQ(), log, auth, db.ParcelsQ(), db.AwaitingParcelsQ(), db.DelayedNotificationsQ()).TakeLater)
			r.Patch("/", parcels.New(db.UsersQ(), log, auth, db.ParcelsQ(), db.AwaitingParcelsQ(), db.DelayedNotificationsQ()).ChangeStatus)
		})

		r.Route("/maintenance", func(r chi.Router) {
			r.Post("/", maintenance.New(db.MaintenanceQ(), db.UsersQ(), log, auth, oneSignal, s3).CreateMaintenance)
			r.Get("/list", maintenance.New(db.MaintenanceQ(), db.UsersQ(), log, auth, oneSignal, s3).GetMaintenanceRequests)
		})
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Glance API"))
	})

	return router
}
