package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type AdminsQ interface {
	Update(admin models.Admin) error
	Insert(admin models.Admin) error
	Delete(admin models.Admin) error
	GetByID(id int64) (models.Admin, error)
}

type AdminsWrapper struct {
	parent *DB
}

func (d *DB) AdminsQ() AdminsQ {
	return &AdminsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (a *AdminsWrapper) Insert(admin models.Admin) error {
	return a.parent.db.Model(&admin).Insert()
}

func (a *AdminsWrapper) Update(admin models.Admin) error {
	return a.parent.db.Model(&admin).Update()
}

func (a *AdminsWrapper) Delete(admin models.Admin) error {
	return a.parent.db.Model(&admin).Delete()
}

func (a *AdminsWrapper) GetByID(id int64) (models.Admin, error) {
	var admin models.Admin
	err := a.parent.db.Select().From(models.AdminsTableName).Where(dbx.HashExp{"id": id}).One(&admin)
	return admin, err
}
