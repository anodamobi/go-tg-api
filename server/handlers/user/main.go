package user

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/sirupsen/logrus"
)

type handler struct {
	db       db.UsersQ
	adminsDB db.AdminsQ
	log      *logrus.Entry
	auth     *config.Authentication
	codeDB   db.CodesForForgottenPwdQ
}

func New(db db.UsersQ, codeDB db.CodesForForgottenPwdQ, adminsDB db.AdminsQ, log *logrus.Entry, jwt *config.Authentication) *handler {
	return &handler{
		db:       db,
		log:      log,
		auth:     jwt,
		codeDB:   codeDB,
		adminsDB: adminsDB,
	}
}
