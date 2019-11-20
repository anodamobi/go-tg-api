package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type NotificationsQ interface {
	Insert(ntf models.Notifications) error
	GetByLocation(continent, country, city, location string) ([]models.Notifications, error)
}

type NotificationsWrapper struct {
	parent *DB
}

func (d *DB) NotificationsQ() NotificationsQ {
	return &NotificationsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (n *NotificationsWrapper) Insert(ntf models.Notifications) error {
	return n.parent.db.Model(&ntf).Insert()
}

func (n *NotificationsWrapper) GetByLocation(continent, country, city, location string) ([]models.Notifications, error) {
	var notifications []models.Notifications
	err := n.parent.db.Select().From(models.NotificationsTableName).
		Where(dbx.HashExp{"continent": continent,
			"country":  country,
			"city":     city,
			"location": location}).
		All(&notifications)
	return notifications, err
}
