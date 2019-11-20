package user

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode login user request")
		errs.BadRequest(w, err)
		return
	}

	user, err := h.db.GetByEmail(request.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).WithField("user_email", request.Email).Error("user does not exist")
			errs.BadRequest(w, errors.New("user does not exist"))
			return
		default:
			h.log.WithError(err).WithField("user_email", request.Email).Error("failed to get user from db")
			errs.InternalError(w)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		h.log.WithError(err).WithField("user_email", request.Email).Error("wrong email or password")
		errs.BadRequest(w, errors.New("wrong email or password"))
		return
	}

	jwtToken, err := h.generateJWT(*user)
	if err != nil {
		h.log.WithError(err).WithField("user_id", user.ID).Error("failed to generate")
		errs.BadRequest(w, errors.New("wrong email or password"))
		return
	}

	response := LoginResponse{Token: jwtToken}

	serializedBody, err := json.Marshal(response)
	if err != nil {
		h.log.WithError(err).WithField("user_id", user.ID).Error("failed to marshal login response")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}

func (h *handler) generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(h.auth.VerifyKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
