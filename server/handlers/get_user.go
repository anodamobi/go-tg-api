package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sirupsen/logrus"

	"github.com/anodamobi/go-tg-api/db"
)

type UserResponse struct {
	ID         string     `json:"id"`
	ExternalID int        `json:"external_id"`
	Name       string     `json:"name"`
	Language   string     `json:"language"`
	Avatar     string     `json:"avatar"`
	JoinedAt   time.Time  `json:"joined_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type UserHandler struct {
	db  *db.DB
	log *logrus.Entry
}

func NewUserHandler(db *db.DB, log *logrus.Entry) *UserHandler {
	return &UserHandler{
		db:  db,
		log: log,
	}
}

func (h UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	//Get user ID
	userID := UserID(r.Context())

	//Get user information from DB
	user, err := h.db.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		h.log.WithField("user_id", userID).
			WithError(err).Error("failed to retrieve user information")
		return
	}

	response, err := json.Marshal(&UserResponse{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Name:       user.Name,
		Language:   user.Language,
		//TODO: return photo link
		Avatar:    user.Avatar,
		JoinedAt:  user.JoinedAt,
		UpdatedAt: &user.UpdatedAt,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		h.log.WithError(err).Error("failed to serialize user information")
		return
	}

	_, _ = w.Write(response)
}

const userIDKey = "id"

//UserID provided from request context, takes ID from JWT claims
func UserID(ctx context.Context) string {
	_, claims, _ := jwtauth.FromContext(ctx)
	userID, ok := claims[userIDKey]
	if !ok {
		return ""
	}

	return userID.(string)
}
