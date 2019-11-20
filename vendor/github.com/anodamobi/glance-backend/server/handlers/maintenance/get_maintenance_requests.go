package maintenance

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
)

func (h *handler) GetMaintenanceRequests(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	requests, err := h.mntncDB.GetByUser(uid)
	if err != nil {
		h.log.WithError(err).Error("failed to get user maintenance request")
		errs.InternalError(w)
		return
	}

	serializedBody, err := json.Marshal(requests)
	if err != nil {
		h.log.WithError(err).Error("failed to marshl get user maintenance request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
