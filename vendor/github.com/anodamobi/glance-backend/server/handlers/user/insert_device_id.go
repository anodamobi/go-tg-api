package user

import (
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) InsertDeviceID(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("device_id")
	if deviceID == "" {
		h.log.Error("device id is empty")
		errs.BadRequest(w, errors.New("device id is empty"))
		return
	}

	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	user, err := h.db.Get(uid)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	user.DeviceID = deviceID

	err = h.db.Update(user)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to update user")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
