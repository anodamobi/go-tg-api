package user

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
)

type GetUserResponse struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Image       string `json:"image"`
	GivenName   string `json:"given_name"`
	Room        int32  `json:"room"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsFirstTime bool   `json:"is_first_time"`
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
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
	resp := GetUserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Image:       user.Image,
		GivenName:   user.GivenName,
		Room:        user.Room,
		Email:       user.Email,
		Phone:       user.Phone,
		IsFirstTime: user.IsFirstTime,
	}

	serializedBody, err := json.Marshal(resp)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to marshal get user request")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
