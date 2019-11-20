package models

const NotificationSettingsTableName = "notification_settings"

type NotificationSettings struct {
	UserID                   int64 `db:"user_id"`
	AllAllowed               bool  `db:"all_allowed"`
	ParcelsAllowed           bool  `db:"parcels_allowed"`
	EventsAllowed            bool  `db:"events_allowed"`
	BuildingBulletinsAllowed bool  `db:"building_bulletins_allowed"`
	MaintenanceAllowed       bool  `db:"maintenance_allowed"`
}

func (n NotificationSettings) TableName() string {
	return NotificationSettingsTableName
}
