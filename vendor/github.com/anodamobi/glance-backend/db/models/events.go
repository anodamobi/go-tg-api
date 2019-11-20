package models

import "time"

const EventsTableName = "events"

const (
	FreeType = iota
	PaidType
)

// types for filters
const (
	Scene = iota
	Balance
	Horizons
	SavedEvents
)

type Event struct {
	ID             int64     `db:"id"`
	Title          string    `db:"title"`
	Date           time.Time `db:"date"`
	Location       string    `db:"location"`
	Category       int32     `db:"category"`
	Type           int32     `db:"type"`
	OpenTo         string    `db:"open_to"`
	MaxAttendees   int32     `db:"max_attendees"`
	Description    *string   `db:"description"`
	IsWaitlistOpen bool      `db:"is_waitlist_open"`
}

func (e Event) TableName() string {
	return EventsTableName
}
