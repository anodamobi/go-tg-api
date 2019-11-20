package db

type NotificationSettingsQ interface {
}

type NotificationSettingsWrapper struct {
	parent *DB
}

func (d *DB) NotificationSettingsQ() NotificationSettingsQ {
	return &NotificationSettingsWrapper{
		parent: &DB{d.db.Clone()},
	}
}
