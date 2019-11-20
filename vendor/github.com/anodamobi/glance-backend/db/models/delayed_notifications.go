package models

import "time"

const DelayedNotificationsTableName = "delayed_notifications"

type DelayedNotifications struct {
	ID           int64     `db:"id"`
	EventID      int64     `db:"event_id"`
	TimeToNotify time.Time `db:"time_to_notify"`
	Comment      string    `db:"comment"`
	Title        string    `db:"title"`
	UserDeviceID string    `db:"user_device_id"`
}

func (d DelayedNotifications) TableName() string {
	return DelayedNotificationsTableName
}
