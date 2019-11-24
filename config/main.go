package config

import (
	"sync"

	"github.com/go-chi/jwtauth"

	"github.com/anodamobi/go-tg-api/db"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	JWT() *jwtauth.JWTAuth
	DB() *db.DB
	Bot() *Bot
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	http *HTTP
	log  *logrus.Entry
	jwt  *jwtauth.JWTAuth
	db   *db.DB
	bot  *Bot
}

func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
