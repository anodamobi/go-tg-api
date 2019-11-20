package user

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type NewAdminReq struct {
	FirstName               string   `json:"first_name"`
	LastName                string   `json:"last_name"`
	Email                   string   `json:"email"`
	Image                   string   `json:"image"`
	SitePermissions         []string `json:"site_permissions,omitempty"`
	AdminPermissions        []string `json:"admin_permissions,omitempty"`
	ResidentPermissions     []string `json:"resident_permissions,omitempty"`
	ChatPermissions         []string `json:"chat_permissions,omitempty"`
	DeliveryPermissions     []string `json:"delivery_permissions,omitempty"`
	MaintenancePermissions  []string `json:"maintenance_permissions,omitempty"`
	WellbeingPermissions    []string `json:"wellbeing_permissions,omitempty"`
	EventsPermissions       []string `json:"events_permissions,omitempty"`
	NotificationPermissions []string `json:"notification_permissions,omitempty"`
}

func (h *handler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var request NewAdminReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode create admin request")
		errs.BadRequest(w, err)
		return
	}

	tempPass := generateTemporaryPass()
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		h.log.WithError(err).Error("failed to generate bytes of temp pass")
		errs.InternalError(w)
		return
	}

	user := models.Admin{
		ID:                      0,
		FirstName:               request.FirstName,
		LastName:                request.LastName,
		Email:                   request.Email,
		Image:                   request.Image,
		Password:                string(passwordBytes),
		SitePermissions:         request.SitePermissions,
		AdminPermissions:        request.AdminPermissions,
		ResidentPermissions:     request.ResidentPermissions,
		ChatPermissions:         request.ChatPermissions,
		DeliveryPermissions:     request.DeliveryPermissions,
		MaintenancePermissions:  request.MaintenancePermissions,
		WellbeingPermissions:    request.WellbeingPermissions,
		EventsPermissions:       request.EventsPermissions,
		NotificationPermissions: request.NotificationPermissions,
	}

	_, err = h.db.GetByEmail(request.Email)
	if err != sql.ErrNoRows {
		h.log.Error("admin already exist")
		errs.BadRequest(w, errors.New("admin already exist"))
		return
	}

	err = h.adminsDB.Insert(user)
	if err != nil {
		h.log.WithError(err).WithField("email", user.Email).Error("failed to insert new admin into db")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
