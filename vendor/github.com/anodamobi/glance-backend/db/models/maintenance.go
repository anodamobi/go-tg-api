package models

import "time"

const MaintenanceTableName = "maintenance"

type Maintenance struct {
	ID          int64     `db:"id" json:"id"`
	Images      []string  `db:"images" json:"images"`
	Description string    `db:"description" json:"description"`
	UserID      int64     `db:"user_id" json:"user_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UniqueID    string    `db:"unique_id" json:"unique_id"`
}

func (m Maintenance) TableName() string {
	return MaintenanceTableName
}
