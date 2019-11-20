package user

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
)

type UpdateUserRequest struct {
	Image     string `json:"image"`
	GivenName string `json:"given_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var request UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode update user request")
		errs.BadRequest(w, err)
		return
	}

	user, err := h.db.Get(uid)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	user = models.User{
		ID:          user.ID,
		Image:       request.Image,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		GivenName:   request.GivenName,
		Room:        user.Room,
		Email:       request.Email,
		Password:    user.Password,
		Phone:       request.Phone,
		IsFirstTime: user.IsFirstTime,
	}

	err = h.db.Update(user)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to update user")
		errs.InternalError(w)
		return
	}

	serializedBody, err := json.Marshal(GetUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Image:     request.Image,
		GivenName: request.GivenName,
		Room:      user.Room,
		Email:     request.Email,
		Phone:     request.Phone,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to marshal updated user")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(serializedBody)
}
