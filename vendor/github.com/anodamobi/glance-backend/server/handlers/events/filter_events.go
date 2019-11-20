package events

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
)

type FilterEventsRequest struct {
	Filters []int32 `json:"filters"`
}

func (h *handler) FilterEvents(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var request FilterEventsRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode filter event request")
		errs.BadRequest(w, err)
		return
	}

	var (
		eventsResp []Event
		eventsTemp []models.Event
	)

	for _, f := range request.Filters {
		events, err := h.db.GetByFilters(f)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
			default:
				h.log.WithError(err).Error("failed to get filtered events from db")
				errs.InternalError(w)
				return
			}
		}

		eventsTemp = append(eventsTemp, events...)
	}

	for _, e := range eventsTemp {
		currentEvent, err := h.userEventsDB.Get(uid, e.ID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				continue
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
			UserStatus:   currentEvent.Status,
			IsSaved:      currentEvent.IsSaved,
		})
	}

	serializedData, err := json.Marshal(GetEventsResponse{Events: eventsResp})
	if err != nil {
		h.log.WithError(err).Error("failed to marshal get filtered events request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedData)
}
