package models

const UsersTableName = "users"

type User struct {
	ID          int64  `db:"id"`
	Image       string `db:"image"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	GivenName   string `db:"given_name"`
	Room        int32  `db:"room"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Phone       string `db:"phone"`
	Continent   string `db:"continent"`
	Country     string `db:"country"`
	City        string `db:"city"`
	Site        string `db:"site"`
	Block       string `db:"block"`
	Floor       int32  `db:"floor"`
	TempPass    string `db:"temp_pass"`
	IsFirstTime bool   `db:"is_first_time"`
	DeviceID    string `db:"device_id"`
}

func (u User) TableName() string {
	return UsersTableName
}
