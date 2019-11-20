package notifications

import (
	"bytes"
	"encoding/json"
	"github.com/anodamobi/glance-backend/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	createNotificationsURL = "https://onesignal.com/api/v1/notifications"
)

type RequestForOneSignal struct {
	AppID             string   `json:"app_id"`
	IncludedSegments  []string `json:"included_segments,omitempty"`
	IncludedPlayerIDs []string `json:"included_player_ids,omitempty"`
	Headings          Text     `json:"headings"`
	Contents          Text     `json:"contents"`
}

type Text struct {
	En string `json:"en,omitempty"`
}

type Notifier struct {
	log       *logrus.Entry
	oneSignal *config.OneSignal
}

func NewNotifier(log *logrus.Entry, oneSignal *config.OneSignal) *Notifier {
	return &Notifier{
		log:       log,
		oneSignal: oneSignal,
	}
}

func (h *Notifier) Notify(subject, message string, userDeviceIDs []string) error {
	serializedBody, err := json.Marshal(RequestForOneSignal{
		AppID:             h.oneSignal.AppID,
		IncludedPlayerIDs: userDeviceIDs,
		Headings:          Text{En: subject},
		Contents:          Text{En: message},
	})
	if err != nil {
		h.log.WithError(err).Error("failed to marshal create notification request for one signal")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, createNotificationsURL, bytes.NewReader(serializedBody))
	if err != nil {
		h.log.WithError(err).Error("failed to create new request for create notification for one signal")
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		h.log.WithError(err).Error("failed to do new request for create notification for one signal")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		h.log.Error("response status from one signal isn't 200")
		return errors.New("response status from one signal isn't 200")
	}

	return nil
}
