package notifications

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

type CreateNotificationRequest struct {
	Type      string `json:"notification_type"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Location  string `json:"location"`
}

func (h *handler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var request CreateNotificationRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to unmarshal create notification request")
		errs.BadRequest(w, err)
		return
	}

	err = h.notifDB.Insert(models.Notifications{
		Type:      request.Type,
		Title:     request.Title,
		Body:      request.Body,
		CreatedAt: time.Now(),
		Sender:    uid,
		Continent: request.Continent,
		Country:   request.Country,
		City:      request.City,
		Location:  request.Location,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert notification to db")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
