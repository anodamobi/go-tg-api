package maintenance

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

type handler struct {
	mntncDB   db.MaintenanceQ
	usersDB   db.UsersQ
	log       *logrus.Entry
	auth      *config.Authentication
	oneSignal *config.OneSignal
	s3        *minio.Client
}

func New(mntncDB db.MaintenanceQ, usersDB db.UsersQ, log *logrus.Entry, jwt *config.Authentication, oneSignal *config.OneSignal, s3 *minio.Client) *handler {
	return &handler{
		mntncDB:   mntncDB,
		oneSignal: oneSignal,
		usersDB:   usersDB,
		log:       log,
		auth:      jwt,
		s3:        s3,
	}
}
