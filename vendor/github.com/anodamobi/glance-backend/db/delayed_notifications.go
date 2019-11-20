package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type DelayedNotificationsQ interface {
	Insert(dn models.DelayedNotifications) error
	GetByIDs(deviceID string, eventID int64) (models.DelayedNotifications, error)
	GetAll() ([]models.DelayedNotifications, error)
}

type DelayedNotificationsWrapper struct {
	parent *DB
}

func (d *DB) DelayedNotificationsQ() DelayedNotificationsQ {
	return &DelayedNotificationsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (d *DelayedNotificationsWrapper) Insert(dn models.DelayedNotifications) error {
	return d.parent.db.Model(&dn).Insert()
}

func (d *DelayedNotificationsWrapper) GetAll() ([]models.DelayedNotifications, error) {
	var dn []models.DelayedNotifications
	err := d.parent.db.Select().From(models.DelayedNotificationsTableName).All(&dn)
	return dn, err
}

func (d *DelayedNotificationsWrapper) GetByIDs(deviceID string, eventID int64) (models.DelayedNotifications, error) {
	var dn models.DelayedNotifications
	err := d.parent.db.Select().From(models.DelayedNotificationsTableName).Where(dbx.HashExp{"event_id": eventID,
		"user_device_id": deviceID}).One(&dn)
	return dn, err
}
