package models

import "time"

const NotificationsTableName = "notifications"

type Notifications struct {
	ID        int64     `db:"id"`
	Type      string    `db:"type"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	Sender    int64     `db:"sender"`
	Continent string    `db:"continent"`
	Country   string    `db:"country"`
	City      string    `db:"city"`
	Location  string    `db:"location"`
}

func (n Notifications) TableName() string {
	return NotificationsTableName
}
