package models

import "time"

const WellbeingRequestsTableName = "wellbeing_requests"

type WellbeingRequest struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	IsResolved bool      `db:"is_resolved"`
}

func (w WellbeingRequest) TableName() string {
	return WellbeingRequestsTableName
}
