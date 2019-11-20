package config

import (
	"github.com/anodamobi/glance-backend/db"
	"github.com/minio/minio-go"
	"sync"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	JWT() *Authentication
	DB() db.QInterface
	OneSignal() *OneSignal
	S3() *minio.Client
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	http      *HTTP
	log       *logrus.Entry
	jwt       *Authentication
	db        db.QInterface
	oneSignal *OneSignal
	s3        *minio.Client
}

func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
