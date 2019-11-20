package parcels

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	usersDB   db.UsersQ
	parcelsDB db.ParcelsQ
	apDB      db.AwaitingParcelsQ
	log       *logrus.Entry
	auth      *config.Authentication
	dnDB      db.DelayedNotificationsQ
}

func New(db db.UsersQ, log *logrus.Entry, jwt *config.Authentication, parcelsDB db.ParcelsQ, apDB db.AwaitingParcelsQ, dnDB db.DelayedNotificationsQ) *handler {
	return &handler{
		usersDB:   db,
		log:       log,
		auth:      jwt,
		parcelsDB: parcelsDB,
		apDB:      apDB,
		dnDB:      dnDB,
	}
}
