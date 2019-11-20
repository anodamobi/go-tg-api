package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type UsersQ interface {
	Insert(user models.User) error
	Get(userID int64) (models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user models.User) error
	GetTempPass() ([]models.User, error)
}

type UsersWrapper struct {
	parent *DB
}

func (d *DB) UsersQ() UsersQ {
	return &UsersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (u *UsersWrapper) Insert(user models.User) error {
	return u.parent.db.Model(&user).Insert()
}

func (u *UsersWrapper) Get(userID int64) (models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"id": userID}).From(models.UsersTableName).One(&user)
	return user, err
}

func (u *UsersWrapper) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"email": email}).From(models.UsersTableName).One(&user)
	return &user, err
}

func (u *UsersWrapper) Update(user models.User) error {
	return u.parent.db.Model(&user).Update()
}

func (u *UsersWrapper) GetTempPass() ([]models.User, error) {
	var user []models.User
	err := u.parent.db.Select().From(models.UsersTableName).All(&user)
	return user, err
}
