package user

import (
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
)

func (h *handler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	adminID := r.URL.Query().Get("aid")
	if adminID == "" {
		h.log.Error("admin id is empty in url")
		errs.BadRequest(w, errors.New("admin id is empty in url"))
		return
	}

	aid, err := cast.ToInt64E(adminID)
	if err != nil {
		h.log.WithError(err).Error("failed to parse admin id to int")
		errs.BadRequest(w, errors.New("admin id isn't integer"))
		return
	}

	err = h.adminsDB.Delete(models.Admin{ID: aid})
	if err != nil {
		h.log.WithError(err).Error("failed to delete admin")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
