package user

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

func (h *handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var request UpdatePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&request)
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
		h.log.WithError(err).WithField("user_id", uid).Error("failed to generate password from source to hash")
		errs.InternalError(w)
		return
	}

	user, err := h.db.Get(uid)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	// TEMPORARY
	splittedEmail := strings.Split(user.Email, "@")
	if len(splittedEmail) > 1 {
		if splittedEmail[1] == "thisisglance.com" {
			user = models.User{
				ID:          user.ID,
				Image:       user.Image,
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				GivenName:   user.GivenName,
				Room:        user.Room,
				Email:       user.Email,
				Password:    string(passwordBytes),
				Phone:       user.Phone,
				IsFirstTime: true,
			}
		} else {
			user = models.User{
				ID:          user.ID,
				Image:       user.Image,
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				GivenName:   user.GivenName,
				Room:        user.Room,
				Email:       user.Email,
				Password:    string(passwordBytes),
				Phone:       user.Phone,
				IsFirstTime: false,
			}
		}
	} else {
		h.log.Error("email doesn't contain @")
		errs.BadRequest(w, errors.New("email doesn't contain @"))
		return
	}

	err = h.db.Update(user)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to update user password")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
