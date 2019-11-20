package userevents

import (
	"database/sql"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
)

func (h *handler) SaveEvent(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	eidRaw := r.URL.Query().Get("event_id")
	if eidRaw == "" {
		h.log.WithError(err).Error("event id is empty")
		errs.BadRequest(w, errors.New("event id is empty"))
		return
	}

	eid, err := cast.ToInt64E(eidRaw)
	if err != nil {
		h.log.WithError(err).Error("failed to cast event id to int64")
		errs.BadRequest(w, err)
		return
	}

	userEvent, err := h.userEventsDB.Get(uid, eid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("event doesn't exist")
			errs.BadRequest(w, errors.New("event doesn't exist"))
			return
		default:
			h.log.WithError(err).Error("failed to get user events")
			errs.InternalError(w)
			return
		}
	}

	userEvent.IsSaved = true

	err = h.userEventsDB.Update(userEvent)
	if err != nil {
		h.log.WithError(err).Error("failed to update user events")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
