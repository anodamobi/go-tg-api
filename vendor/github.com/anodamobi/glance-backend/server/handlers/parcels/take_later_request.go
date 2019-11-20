package parcels

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

type TakeLaterReq struct {
	ParcelID int64     `json:"parcel_id"`
	Date     time.Time `json:"date"`
}

func (h *handler) TakeLater(w http.ResponseWriter, r *http.Request) {
	var req TakeLaterReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse take parcel later request")
		errs.BadRequest(w, err)
		return
	}

	parcel, err := h.parcelsDB.Get(req.ParcelID)
	if err != nil {
		h.log.WithError(err).WithField("parcel_id", req.ParcelID).Error("failed to get parcel from db")
		errs.InternalError(w)
		return
	}

	parcel.Status = models.TakeLater

	err = h.parcelsDB.Update(parcel)
	if err != nil {
		h.log.WithError(err).WithField("parcel_id", req.ParcelID).Error("failed to update parcel")
		errs.InternalError(w)
		return
	}

	err = h.apDB.Insert(models.AwaitingParcel{
		PID:      req.ParcelID,
		TakeTime: req.Date,
	})
	if err != nil {
		h.log.WithError(err).WithField("parcel_id", req.ParcelID).Error("failed to create new take later parcel request")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
