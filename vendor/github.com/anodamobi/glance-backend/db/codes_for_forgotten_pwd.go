package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type CodesForForgottenPwdQ interface {
	Insert(code models.CodesForForgottenPwd) error
	Get(email, cd string) (models.CodesForForgottenPwd, error)
	GetAll() ([]models.CodesForForgottenPwd, error)
}

type CodesForForgottenPwdWrapper struct {
	parent *DB
}

func (d *DB) CodesForForgottenPwdQ() CodesForForgottenPwdQ {
	return &CodesForForgottenPwdWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (c *CodesForForgottenPwdWrapper) Insert(code models.CodesForForgottenPwd) error {
	return c.parent.db.Model(&code).Insert()
}

func (c *CodesForForgottenPwdWrapper) Get(email, cd string) (models.CodesForForgottenPwd, error) {
	var codes models.CodesForForgottenPwd
	err := c.parent.db.Select().From(models.CodesForForgottenPwdTableName).Where(dbx.HashExp{"email": email, "code": cd}).One(&codes)
	return codes, err
}

func (c *CodesForForgottenPwdWrapper) GetAll() ([]models.CodesForForgottenPwd, error) {
	var codes []models.CodesForForgottenPwd
	err := c.parent.db.Select().From(models.CodesForForgottenPwdTableName).All(&codes)
	return codes, err
}
