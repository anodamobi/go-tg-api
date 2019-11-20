package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type UserEventsQ interface {
	Insert(ue models.UserEvents) error
	Get(uid, eid int64) (models.UserEvents, error)
	Update(ue models.UserEvents) error
	GetEvent(eid int64) ([]models.UserEvents, error)
}

type UserEventsWrapper struct {
	parent *DB
}

func (d *DB) UserEventsQ() UserEventsQ {
	return &UserEventsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (u *UserEventsWrapper) Insert(ue models.UserEvents) error {
	return u.parent.db.Model(&ue).Insert()
}

func (u *UserEventsWrapper) Update(ue models.UserEvents) error {
	return u.parent.db.Model(&ue).Update()
}

func (u *UserEventsWrapper) Get(uid, eid int64) (models.UserEvents, error) {
	var ue models.UserEvents
	err := u.parent.db.Select().From(models.UserEventsTableName).Where(dbx.HashExp{"user_id": uid, "event_id": eid}).One(&ue)
	return ue, err
}

func (u *UserEventsWrapper) GetEvent(eid int64) ([]models.UserEvents, error) {
	var ue []models.UserEvents
	err := u.parent.db.Select().From(models.UserEventsTableName).Where(dbx.HashExp{"event_id": eid}).All(&ue)
	return ue, err
}
