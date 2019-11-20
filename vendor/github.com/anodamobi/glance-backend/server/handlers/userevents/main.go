package userevents

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	usersDB      db.UsersQ
	userEventsDB db.UserEventsQ
	log          *logrus.Entry
	auth         *config.Authentication
	oneSignal    *config.OneSignal
	dnDB         db.DelayedNotificationsQ
	eventDB      db.EventsQ
}

func New(db db.UsersQ, log *logrus.Entry, jwt *config.Authentication, userEventsDB db.UserEventsQ,
	oneSignal *config.OneSignal, dnDB db.DelayedNotificationsQ, eventDB db.EventsQ) *handler {
	return &handler{
		usersDB:      db,
		log:          log,
		auth:         jwt,
		userEventsDB: userEventsDB,
		oneSignal:    oneSignal,
		dnDB:         dnDB,
		eventDB:      eventDB,
	}
}
