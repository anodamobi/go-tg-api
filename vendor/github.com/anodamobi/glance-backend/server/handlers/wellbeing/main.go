package wellbeing

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	usersDB db.UsersQ
	wrDB    db.WellbeingRequestsQ
	log     *logrus.Entry
	auth    *config.Authentication
}

func New(db db.UsersQ, log *logrus.Entry, jwt *config.Authentication, wrDB db.WellbeingRequestsQ) *handler {
	return &handler{
		usersDB: db,
		log:     log,
		auth:    jwt,
		wrDB:    wrDB,
	}
}
