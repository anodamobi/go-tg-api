package models

import "time"

const ParcelsTableName = "parcels"

type Parcel struct {
	ID                int64     `db:"id"`
	UserID            int64     `db:"user_id"`
	Room              int32     `db:"room"`
	DeliveryCompany   string    `db:"delivery_company"`
	ParcelLocation    string    `db:"parcel_location"`
	NotificationTitle string    `db:"notification_title"`
	NotificationText  string    `db:"notification_text"`
	CreatedAt         time.Time `db:"created_at"`
	Status            int32     `db:"status"`
}

const (
	NotCollected = iota
	TakeNow
	TakeLater
	LeaveInRoom
)

func (p Parcel) TableName() string {
	return ParcelsTableName
}
