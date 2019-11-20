package events

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	db           db.EventsQ
	userEventsDB db.UserEventsQ
	log          *logrus.Entry
	auth         *config.Authentication
}

func New(db db.EventsQ, log *logrus.Entry, jwt *config.Authentication, userEventsDB db.UserEventsQ) *handler {
	return &handler{
		db:           db,
		log:          log,
		auth:         jwt,
		userEventsDB: userEventsDB,
	}
}
