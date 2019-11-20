package userevents

import (
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/auth"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/pkg/errors"
	"net/http"
)

type ChangeStatusRequest struct {
	Status  int   `json:"status"`
	EventID int64 `json:"event_id"`
}

func (h *handler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	uid, _, err := auth.GetIDFromJWT(r, h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from jwt")
		errs.BadRequest(w, err)
		return
	}

	var req ChangeStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse change status request")
		errs.BadRequest(w, err)
		return
	}

	//user, err := h.usersDB.Get(uid)
	//if err != nil {
	//	h.log.WithError(err).Error("failed to get user from db")
	//	errs.InternalError(w)
	//	return
	//}

	userEvent, err := h.userEventsDB.Get(uid, req.EventID)
	if err != nil {
		h.log.WithError(err).Error("failed to get user events")
		errs.InternalError(w)
		return
	}

	//event, err := h.eventDB.GetByID(req.EventID)
	//if err != nil {
	//	h.log.WithError(err).Error("failed to get event")
	//	errs.InternalError(w)
	//	return
	//}

	var parsedStatus int
	switch req.Status {
	case 0:
		parsedStatus = models.Canceled
		//err := notifications.NewNotifier(h.log, h.oneSignal).Notify("You are out!", "You have been cancelled booked event!", []string{user.DeviceID})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create push notification for cancel event")
		//	errs.InternalError(w)
		//	return
		//}
	case 1:
		parsedStatus = models.Joined
		//err := notifications.NewNotifier(h.log, h.oneSignal).Notify("You are in!", "You have been booked event!", []string{user.DeviceID})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create push notification for booking event")
		//	errs.InternalError(w)
		//	return
		//}
		//
		//year, month, day := event.Date.Date()
		//timeToNotify := time.Date(year, month, day, 9, 0, 0, 0, time.UTC)
		//
		//err = h.dnDB.Insert(models.DelayedNotifications{
		//	EventID:      req.EventID,
		//	TimeToNotify: timeToNotify,
		//	Comment:      "You have been attended to event this evening!",
		//	Title:        "Upcoming event",
		//	UserDeviceID: user.DeviceID,
		//})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create delayed notification about user upcoming event")
		//	errs.InternalError(w)
		//	return
		//}
		//
		//timeToNotify = time.Date(year, month, day-1, 9, 0, 0, 0, time.UTC)
		//err = h.dnDB.Insert(models.DelayedNotifications{
		//	EventID:      req.EventID,
		//	TimeToNotify: timeToNotify,
		//	Comment:      "You have been attended to event next day!",
		//	Title:        "Upcoming event",
		//	UserDeviceID: user.DeviceID,
		//})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create delayed notification about user next day event")
		//	errs.InternalError(w)
		//	return
		//}
		//
		//timeToNotify = time.Date(year, month, day-3, 23, 59, 0, 0, time.UTC)
		//err = h.dnDB.Insert(models.DelayedNotifications{
		//	EventID:      req.EventID,
		//	TimeToNotify: timeToNotify,
		//	Comment:      "You have been attended to event in 3 days!",
		//	Title:        "Upcoming event",
		//	UserDeviceID: user.DeviceID,
		//})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create delayed notification about user 3 days event")
		//	errs.InternalError(w)
		//	return
		//}

	case 2:
		parsedStatus = models.Pending
	case 3:
		parsedStatus = models.Waitlist
	default:
		h.log.Error("invalid status")
		errs.BadRequest(w, errors.New("invalid status"))
		return
	}

	userEvent.Status = int32(parsedStatus)

	err = h.userEventsDB.Update(userEvent)
	if err != nil {
		h.log.WithError(err).Error("failed to update user events")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
