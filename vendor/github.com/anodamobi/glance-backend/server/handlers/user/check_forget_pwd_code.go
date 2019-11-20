package user

import (
	"database/sql"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
)

func (h *handler) CheckForgetPwdCode(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		h.log.Error("email is empty in url")
		errs.BadRequest(w, errors.New("email is empty in url"))
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		h.log.Error("code is empty in url")
		errs.BadRequest(w, errors.New("code is empty in url"))
		return
	}

	_, err := h.codeDB.Get(email, code)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.Error("wrong code")
			errs.BadRequest(w, errors.New("wrong code"))
			return
		default:
			h.log.WithError(err).Error("failed to get code from db")
			errs.InternalError(w)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
