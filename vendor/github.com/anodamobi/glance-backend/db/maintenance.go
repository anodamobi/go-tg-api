package db

import (
	"github.com/anodamobi/glance-backend/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type MaintenanceQ interface {
	Insert(maintenance models.Maintenance) error
	GetByUser(userID int64) ([]models.Maintenance, error)
}

type MaintenanceWrapper struct {
	parent *DB
}

func (d *DB) MaintenanceQ() MaintenanceQ {
	return &MaintenanceWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (m *MaintenanceWrapper) Insert(maintenance models.Maintenance) error {
	return m.parent.db.Model(&maintenance).Insert()
}

func (m *MaintenanceWrapper) GetByUser(userID int64) ([]models.Maintenance, error) {
	var maintenances []models.Maintenance
	err := m.parent.db.Select().Where(dbx.HashExp{"user_id": userID}).From(models.MaintenanceTableName).All(&maintenances)
	return maintenances, err
}
