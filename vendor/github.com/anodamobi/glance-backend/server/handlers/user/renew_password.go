package user

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RenewPasswordRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *handler) RenewPassword(w http.ResponseWriter, r *http.Request) {
	var request RenewPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode update user password request")
		errs.BadRequest(w, err)
		return
	}

	if request.Password == "" {
		h.log.Error("empty password in body")
		errs.BadRequest(w, errors.New("empty password in body"))
		return
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		h.log.WithError(err).Error("failed to generate password from source to hash")
		errs.InternalError(w)
		return
	}

	user, err := h.db.GetByEmail(request.Email)
	if err != nil {
		h.log.WithError(err).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	user.Password = string(passwordBytes)

	err = h.db.Update(*user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user password")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
