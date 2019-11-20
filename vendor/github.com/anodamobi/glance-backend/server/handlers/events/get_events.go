package events

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

type GetEventsResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	Category     int32     `json:"category"`
	Type         int32     `json:"payment_type"`
	OpenTo       string    `json:"open_to"`
	MaxAttendees int32     `json:"max_attendees"`
	IsSaved      bool      `json:"is_saved"`
	Description  *string   `json:"description,omitempty"`
	UserStatus   int32     `json:"user_status"`
}

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.db.Get()
	if err != nil {
		h.log.WithError(err).Error("failed to get all events from db")
		errs.InternalError(w)
		return
	}

	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var eventsResp []Event
	for _, e := range events {
		currentEvent, err := h.userEventsDB.Get(uid, e.ID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
			default:
				h.log.WithError(err).WithField("event_id", e.ID).Error("failed to get user's current event")
				errs.InternalError(w)
				return
			}
		}

		eventsResp = append(eventsResp, Event{
			Title:        e.Title,
			Date:         e.Date,
			Location:     e.Location,
			Category:     e.Category,
			Type:         e.Type,
			OpenTo:       e.OpenTo,
			MaxAttendees: e.MaxAttendees,
			Description:  e.Description,
			ID:           e.ID,
			IsSaved:      currentEvent.IsSaved,
			UserStatus:   currentEvent.Status,
		})
	}

	serializedBody, err := json.Marshal(GetEventsResponse{Events: eventsResp})
	if err != nil {
		h.log.WithError(err).Error("failed to marshal get events request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
