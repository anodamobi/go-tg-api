package parcels

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type ChangeStatusReq struct {
	ParcelID  int64 `json:"parcel_id"`
	NewStatus int32 `json:"new_status"`
}

func (h *handler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	var req ChangeStatusReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse change parcel status request")
		errs.BadRequest(w, err)
		return
	}

	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	user, err := h.usersDB.Get(uid)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	parcel, err := h.parcelsDB.Get(req.ParcelID)
	if err != nil {
		h.log.WithError(err).WithField("parcel_id", req.ParcelID).Error("failed to get parcel from db")
		errs.InternalError(w)
		return
	}

	switch req.NewStatus {
	case models.LeaveInRoom:
		parcel.Status = models.LeaveInRoom
	case models.TakeNow:
		parcel.Status = models.TakeNow

		year, month, day := time.Now().Date()
		timeToNotify := time.Date(year, month, day+1, time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)

		err = h.dnDB.Insert(models.DelayedNotifications{
			EventID:      0,
			TimeToNotify: timeToNotify,
			Comment:      "Your parcel is holding on the reception and waiting you!",
			Title:        "Take your parcel",
			UserDeviceID: user.DeviceID,
		})
		if err != nil {
			h.log.WithError(err).WithField("parcel_id", req.ParcelID).
				Error("failed to insert new delayed notification about parcel which is on reception for 1 day")
			errs.InternalError(w)
			return
		}
	default:
		h.log.Error("wrong status number")
		errs.BadRequest(w, errors.New("wrong status number"))
		return
	}

	err = h.parcelsDB.Update(parcel)
	if err != nil {
		h.log.WithError(err).WithField("parcel_id", req.ParcelID).Error("failed to update parcel with new status")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
