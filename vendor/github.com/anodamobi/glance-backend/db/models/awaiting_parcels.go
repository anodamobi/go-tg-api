package models

import "time"

const AwaitingParcelsTableName = "awaiting_parcels"

type AwaitingParcel struct {
	PID      int64     `db:"pid"`
	TakeTime time.Time `db:"take_time"`
}

func (a AwaitingParcel) TableName() string {
	return AwaitingParcelsTableName
}
