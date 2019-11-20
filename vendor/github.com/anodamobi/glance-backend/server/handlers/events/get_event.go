package events

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

type GetEventResponse struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Date           time.Time `json:"date"`
	Location       string    `json:"location"`
	Category       int32     `json:"category"`
	Type           int32     `json:"payment_type"`
	OpenTo         string    `json:"open_to"`
	MaxAttendees   int32     `json:"max_attendees"`
	Description    *string   `json:"description"`
	IsSaved        bool      `json:"is_saved"`
	IsWaitlistOpen bool      `json:"is_waitlist_open"`
	Status         int32     `json:"status"`
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eidRaw := r.URL.Query().Get("eid")
	if eidRaw == "" {
		h.log.Error("event id is empty")
		errs.BadRequest(w, errors.New("event id is empty"))
		return
	}

	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	eid, err := cast.ToInt64E(eidRaw)
	if err != nil {
		h.log.WithError(err).Error("failed to cast event id to int64")
		errs.BadRequest(w, err)
		return
	}

	event, err := h.db.GetByID(eid)
	if err != nil {
		h.log.WithError(err).WithField("event_id", eid).Error("failed to get event from db")
		errs.InternalError(w)
		return
	}

	userEvents, err := h.userEventsDB.GetEvent(eid)
	if err != nil {
		h.log.WithError(err).WithField("event_id", eid).Error("failed to get user's event from db")
		errs.InternalError(w)
		return
	}

	attendees := len(userEvents)

	if int32(attendees) >= event.MaxAttendees && !event.IsWaitlistOpen {
		event.IsWaitlistOpen = true

		err := h.db.Update(event)
		if err != nil {
			h.log.WithError(err).WithField("event_id", eid).Error("failed to get update event with ON waitlist")
			errs.InternalError(w)
			return
		}
	}

	currentEvent, err := h.userEventsDB.Get(uid, event.ID)
	if err != nil {
		h.log.WithError(err).WithField("event_id", eid).Error("failed to get user's current event")
		errs.InternalError(w)
		return
	}

	response := GetEventResponse{
		ID:             event.ID,
		Title:          event.Title,
		Date:           event.Date,
		Location:       event.Location,
		Category:       event.Category,
		Type:           event.Type,
		OpenTo:         event.OpenTo,
		MaxAttendees:   event.MaxAttendees,
		Description:    event.Description,
		IsWaitlistOpen: event.IsWaitlistOpen,
		IsSaved:        currentEvent.IsSaved,
		Status:         currentEvent.Status,
	}

	serializedBody, err := json.Marshal(response)
	if err != nil {
		h.log.WithError(err).Error("failed to marshal get event request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
