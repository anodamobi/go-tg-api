package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type EventsQ interface {
	Insert(event models.Event) error
	Get() ([]models.Event, error)
	GetByFilters(filter int32) ([]models.Event, error)
	GetByID(eid int64) (models.Event, error)
	Update(event models.Event) error
}

type EventsWrapper struct {
	parent *DB
}

func (d *DB) EventsQ() EventsQ {
	return &EventsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (e *EventsWrapper) Insert(event models.Event) error {
	return e.parent.db.Model(&event).Insert()
}

func (e *EventsWrapper) Get() ([]models.Event, error) {
	var events []models.Event
	err := e.parent.db.Select().From(models.EventsTableName).All(&events)
	return events, err
}

func (e *EventsWrapper) GetByFilters(filter int32) ([]models.Event, error) {
	var events []models.Event
	err := e.parent.db.Select().From(models.EventsTableName).Where(dbx.HashExp{"category": filter}).All(&events)
	return events, err
}

func (e *EventsWrapper) GetByID(eid int64) (models.Event, error) {
	var event models.Event
	err := e.parent.db.Select().From(models.EventsTableName).Where(dbx.HashExp{"id": eid}).One(&event)
	return event, err
}

func (e *EventsWrapper) Update(event models.Event) error {
	return e.parent.db.Model(&event).Update()
}
