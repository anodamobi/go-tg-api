package user

import (
	"crypto/rand"
	"database/sql"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) SendCodeForRenewPassword(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		h.log.Error("email is empty in url")
		errs.BadRequest(w, errors.New("email is empty in url"))
		return
	}

	_, err := h.db.GetByEmail(email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.Error("user with given email doesn't exist")
			errs.BadRequest(w, errors.New("user with given email doesn't exist"))
			return
		default:
			h.log.WithError(err).Error("failed to get user from database")
			errs.InternalError(w)
			return
		}
	}

	code := code()

	err = h.codeDB.Insert(models.CodesForForgottenPwd{
		Code:  code,
		Email: email,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert new code to db")
		errs.InternalError(w)
		return
	}

	// todo: create emailing via zoho for send code to the user

	w.WriteHeader(http.StatusOK)
}

func code() string {
	b := make([]byte, 40)
	_, _ = rand.Read(b)

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, a := range b {
		b[i] = letters[a%byte(len(letters))]
	}

	return string(b[:5])
}
