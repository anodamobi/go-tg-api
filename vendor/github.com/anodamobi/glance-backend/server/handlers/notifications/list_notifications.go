package notifications

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"net/http"
	"time"
)

type ListNotificationResp struct {
	List []Notification `json:"list"`
}

type Notification struct {
	ID               int64     `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	Title            string    `json:"title"`
	NotificationType string    `json:"notification_type"`
	Sender           string    `json:"sender"`
}

func (h *handler) ListNotifications(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	user, err := h.usersDB.Get(uid)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	userNotifications, err := h.notifDB.GetByLocation(user.Continent, user.Country, user.City, user.Block)
	if err != nil {
		h.log.WithError(err).WithField("user_id", uid).Error("failed to get user notifications from db")
		errs.InternalError(w)
		return
	}

	var notifications []Notification
	for _, n := range userNotifications {
		sender, err := h.usersDB.Get(n.Sender)
		if err != nil {
			h.log.WithError(err).WithField("sender_id", n.Sender).Error("failed to get notification sender user profile")
			errs.InternalError(w)
			return
		}

		notifications = append(notifications, Notification{
			ID:               n.ID,
			CreatedAt:        n.CreatedAt,
			Title:            n.Title,
			NotificationType: n.Type,
			Sender:           sender.FirstName,
		})
	}

	serializedBody, err := json.Marshal(ListNotificationResp{List: notifications})
	if err != nil {
		h.log.WithError(err).Error("failed to marshal notification list response")
		errs.InternalError(w)
		return
	}

	_, _ = w.Write(serializedBody)
}
