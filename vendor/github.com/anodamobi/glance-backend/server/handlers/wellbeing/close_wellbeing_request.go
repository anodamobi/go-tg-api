package wellbeing

import (
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
)

func (h *handler) CloseWellbeingRequest(w http.ResponseWriter, r *http.Request) {
	rid := r.URL.Query().Get("request_id")
	if rid == "" {
		h.log.Error("request id empty in url")
		errs.BadRequest(w, errors.New("request id empty in url"))
		return
	}

	id, err := cast.ToInt64E(rid)
	if err != nil {
		h.log.WithError(err).Error("failed to cast request id to int64")
		errs.BadRequest(w, err)
		return
	}

	wr, err := h.wrDB.Get(id)
	if err != nil {
		h.log.WithError(err).WithField("request_id", id).Error("failed to get wellbeing request")
		errs.InternalError(w)
		return
	}

	wr.IsResolved = true

	err = h.wrDB.Update(wr)
	if err != nil {
		h.log.WithError(err).WithField("request_id", id).Error("failed to update wellbeing request to done status")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
