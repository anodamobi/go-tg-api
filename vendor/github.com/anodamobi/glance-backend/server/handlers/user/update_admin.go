package user

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
)

type UpdateAdminRequest struct {
	AdminID                 int64    `json:"admin_id"`
	FirstName               string   `db:"first_name"`
	LastName                string   `db:"last_name"`
	Email                   string   `db:"email"`
	Image                   string   `db:"image"`
	SitePermissions         []string `db:"site_permissions"`
	AdminPermissions        []string `db:"admin_permissions"`
	ResidentPermissions     []string `db:"resident_permissions"`
	ChatPermissions         []string `db:"chat_permissions"`
	DeliveryPermissions     []string `db:"delivery_permissions"`
	MaintenancePermissions  []string `db:"maintenance_permissions"`
	WellbeingPermissions    []string `db:"wellbeing_permissions"`
	EventsPermissions       []string `db:"events_permissions"`
	NotificationPermissions []string `db:"notification_permissions"`
}

func (h *handler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var req UpdateAdminRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to decode update admin request")
		errs.BadRequest(w, err)
		return
	}

	admin, err := h.adminsDB.GetByID(req.AdminID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("admin with given id doesnt exist")
			errs.BadRequest(w, errors.New("admin with given id doesnt exist"))
			return
		default:
			h.log.WithError(err).Error("failed to get admin from db")
			errs.InternalError(w)
			return
		}
	}

	admin = models.Admin{
		ID:                      admin.ID,
		FirstName:               req.FirstName,
		LastName:                req.LastName,
		Email:                   req.Email,
		Image:                   req.Image,
		Password:                admin.Password,
		SitePermissions:         req.SitePermissions,
		AdminPermissions:        req.AdminPermissions,
		ResidentPermissions:     req.ResidentPermissions,
		ChatPermissions:         req.ChatPermissions,
		DeliveryPermissions:     req.DeliveryPermissions,
		MaintenancePermissions:  req.MaintenancePermissions,
		WellbeingPermissions:    req.WellbeingPermissions,
		EventsPermissions:       req.EventsPermissions,
		NotificationPermissions: req.NotificationPermissions,
	}

	err = h.adminsDB.Update(admin)
	if err != nil {
		h.log.WithError(err).Error("failed to update admin from db")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
