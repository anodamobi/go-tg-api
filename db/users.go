package db

import (
	"time"

	"github.com/go-ozzo/ozzo-dbx"
)

const usersTableName = "users"

type User struct {
	ID         string    `db:"id"`
	ExternalID int       `db:"external_id"`
	Name       string    `db:"name"`
	Language   string    `db:"language"`
	Avatar     string    `db:"avatar"`
	JoinedAt   time.Time `db:"joined_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (u User) TableName() string {
	return usersTableName
}

func (d *DB) CreateUser(user *User) error {
	return d.db.Model(user).Insert()
}

func (d *DB) GetUser(id string) (*User, error) {
	wallet := &User{}
	err := d.db.Select().
		Where(dbx.HashExp{"id": id}).
		One(wallet)
	return wallet, err
}
