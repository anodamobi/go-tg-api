package config

import (
	"sync"

	"github.com/anodamobi/go-tg-api/db"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	JWT() *Authentication
	DB() *db.DB
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	http *HTTP
	log  *logrus.Entry
	jwt  *Authentication
	db   *db.DB
}

func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
