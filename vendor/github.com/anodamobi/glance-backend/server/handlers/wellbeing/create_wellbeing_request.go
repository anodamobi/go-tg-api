package wellbeing

import (
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

func (h *handler) CreateWellbeingRequest(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	err = h.wrDB.Insert(models.WellbeingRequest{
		UserID:     uid,
		CreatedAt:  time.Now(),
		IsResolved: false,
	})
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to create user wellbeing requests")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
