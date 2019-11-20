package models

const UserEventsTableName = "user_events"

type UserEvents struct {
	ID      int64 `db:"id"`
	EventID int64 `db:"event_id"`
	UserID  int64 `db:"user_id"`
	IsSaved bool  `db:"is_saved"`
	Status  int32 `db:"status"`
}

func (u UserEvents) TableName() string {
	return UserEventsTableName
}

const (
	Canceled = iota
	Joined
	Pending
	Waitlist
)
