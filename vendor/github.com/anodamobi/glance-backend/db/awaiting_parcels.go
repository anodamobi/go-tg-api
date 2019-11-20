package db

import "github.com/anodamobi/glance-backend/db/models"

type AwaitingParcelsQ interface {
	Insert(parcel models.AwaitingParcel) error
}

type AwaitingParcelsWrapper struct {
	parent *DB
}

func (d *DB) AwaitingParcelsQ() AwaitingParcelsQ {
	return &AwaitingParcelsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (p *AwaitingParcelsWrapper) Insert(parcel models.AwaitingParcel) error {
	return p.parent.db.Model(&parcel).Insert()
}
