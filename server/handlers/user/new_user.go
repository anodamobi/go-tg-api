package user

import (
	"database/sql"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

type NewUserByAdminRequest struct {
	Email     string `json:"email"`
	Room      int32  `json:"room"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Site      string `json:"site"`
	Block     string `json:"block"`
	Floor     int32  `json:"floor"`
}

func (h *handler) NewUserByAdmin(w http.ResponseWriter, r *http.Request) {
	var request NewUserByAdminRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.WithError(err).Error("failed to decode new user by admin request")
		errs.BadRequest(w, err)
		return
	}

	tempPass := generateTemporaryPass()
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		h.log.WithError(err).Error("failed to generate bytes of temp pass")
		errs.InternalError(w)
		return
	}

	user := models.User{
		Image:       "",
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		GivenName:   "",
		Room:        request.Room,
		Email:       request.Email,
		Password:    string(passwordBytes),
		Phone:       "",
		Continent:   request.Continent,
		Country:     request.Country,
		City:        request.City,
		Site:        request.Site,
		Block:       request.Block,
		Floor:       request.Floor,
		TempPass:    tempPass,
		IsFirstTime: true,
		DeviceID:    "",
	}

	_, err = h.db.GetByEmail(request.Email)
	if err != sql.ErrNoRows {
		h.log.Error("user already exist")
		errs.BadRequest(w, errors.New("user already exist"))
		return
	}

	err = h.db.Insert(user)
	if err != nil {
		h.log.WithError(err).WithField("email", user.Email).Error("failed to insert new user into db")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func generateTemporaryPass() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
