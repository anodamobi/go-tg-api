package notifications

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	notifDB   db.NotificationsQ
	usersDB   db.UsersQ
	auth      *config.Authentication
	log       *logrus.Entry
	oneSignal *config.OneSignal
}

func New(notifDB db.NotificationsQ, usersDB db.UsersQ, log *logrus.Entry, jwt *config.Authentication, oneSignal *config.OneSignal) *handler {
	return &handler{
		notifDB:   notifDB,
		usersDB:   usersDB,
		log:       log,
		auth:      jwt,
		oneSignal: oneSignal,
	}
}
