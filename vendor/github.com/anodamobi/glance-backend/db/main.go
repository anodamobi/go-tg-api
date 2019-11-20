package db

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

//go:generate mockery -all

// QInterface represent interface of the access to the user db
type QInterface interface {
	UsersQ() UsersQ
	EventsQ() EventsQ
	UserEventsQ() UserEventsQ
	NotificationsQ() NotificationsQ
	WellbeingRequestsQ() WellbeingRequestsQ
	ParcelsQ() ParcelsQ
	AwaitingParcelsQ() AwaitingParcelsQ
	MaintenanceQ() MaintenanceQ
	DelayedNotificationsQ() DelayedNotificationsQ
	CodesForForgottenPwdQ() CodesForForgottenPwdQ
	AdminsQ() AdminsQ

	DBX() *dbx.DB
}

type DB struct {
	db *dbx.DB
}

func (d DB) DBX() *dbx.DB {
	return d.db
}

func New(link string) (QInterface, error) {
	db, err := dbx.Open("postgres", link)
	return &DB{db: db}, err
}
