package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type ParcelsQ interface {
	Insert(parcel models.Parcel) error
	Get(pid int64) (models.Parcel, error)
	Update(parcel models.Parcel) error
}

type ParcelsWrapper struct {
	parent *DB
}

func (d *DB) ParcelsQ() ParcelsQ {
	return &ParcelsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (p *ParcelsWrapper) Insert(parcel models.Parcel) error {
	return p.parent.db.Model(&parcel).Insert()
}

func (p *ParcelsWrapper) Get(pid int64) (models.Parcel, error) {
	var parcel models.Parcel
	err := p.parent.db.Select().Where(dbx.HashExp{"id": pid}).From(models.ParcelsTableName).One(&parcel)
	return parcel, err
}

func (p *ParcelsWrapper) Update(parcel models.Parcel) error {
	return p.parent.db.Model(&parcel).Update()
}
