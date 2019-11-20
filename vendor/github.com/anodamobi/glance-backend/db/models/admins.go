package models

const AdminsTableName = "admins"

type Admin struct {
	ID                      int64    `db:"id"`
	FirstName               string   `db:"first_name"`
	LastName                string   `db:"last_name"`
	Email                   string   `db:"email"`
	Image                   string   `db:"image"`
	Password                string   `db:"password"`
	SitePermissions         []string `db:"site_permissions"`
	AdminPermissions        []string `db:"admin_permissions"`
	ResidentPermissions     []string `db:"resident_permissions"`
	ChatPermissions         []string `db:"chat_permissions"`
	DeliveryPermissions     []string `db:"delivery_permissions"`
	MaintenancePermissions  []string `db:"maintenance_permissions"`
	WellbeingPermissions    []string `db:"wellbeing_permissions"`
	EventsPermissions       []string `db:"events_permissions"`
	NotificationPermissions []string `db:"notification_permissions"`
}

func (Admin) TableName() string {
	return AdminsTableName
}
