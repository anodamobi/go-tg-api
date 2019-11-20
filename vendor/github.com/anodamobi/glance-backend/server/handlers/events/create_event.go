package events

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

type CreateEventRequest struct {
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	Category     int32     `json:"category"`
	OpenTo       string    `json:"open_to"`
	MaxAttendees int32     `json:"max_attendees"`
	Description  *string   `json:"description,omitempty"`
}

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var request CreateEventRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode create event request")
		errs.BadRequest(w, err)
		return
	}

	event := models.Event{
		Title:        request.Title,
		Date:         request.Date,
		Location:     request.Location,
		Category:     request.Category,
		Type:         models.FreeType,
		OpenTo:       request.OpenTo,
		MaxAttendees: request.MaxAttendees,
		Description:  request.Description,
	}

	err = h.db.Insert(event)
	if err != nil {
		h.log.WithError(err).WithField("event_title", request.Title).Error("failed to input new event into db")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
