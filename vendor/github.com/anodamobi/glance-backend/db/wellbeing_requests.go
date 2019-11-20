package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type WellbeingRequestsQ interface {
	Insert(req models.WellbeingRequest) error
	Get(id int64) (models.WellbeingRequest, error)
	Update(req models.WellbeingRequest) error
}

type WellbeingRequestsWrapper struct {
	parent *DB
}

func (d *DB) WellbeingRequestsQ() WellbeingRequestsQ {
	return &WellbeingRequestsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (w *WellbeingRequestsWrapper) Insert(req models.WellbeingRequest) error {
	return w.parent.db.Model(&req).Insert()
}

func (w *WellbeingRequestsWrapper) Get(id int64) (models.WellbeingRequest, error) {
	var request models.WellbeingRequest
	err := w.parent.db.Select().From(models.WellbeingRequestsTableName).Where(dbx.HashExp{"id": id}).One(&request)
	return request, err
}

func (w *WellbeingRequestsWrapper) Update(req models.WellbeingRequest) error {
	return w.parent.db.Model(&req).Update()
}
