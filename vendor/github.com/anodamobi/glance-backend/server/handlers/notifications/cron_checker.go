package notifications

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/anodamobi/glance-backend/db"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"time"
)

type EventNotification struct {
	delNotDB  db.DelayedNotificationsQ
	log       *logrus.Entry
	oneSignal *config.OneSignal
}

func NewEventNotificationWorker(delNotDB db.DelayedNotificationsQ, log *logrus.Entry, oneSignal *config.OneSignal) EventNotification {
	return EventNotification{
		delNotDB:  delNotDB,
		log:       log,
		oneSignal: oneSignal,
	}
}

func (e EventNotification) Check() {
	c := cron.New()
	_, err := c.AddFunc("@hourly", func() {
		allDelayedNotifications, err := e.delNotDB.GetAll()
		if err != nil {
			e.log.WithError(err).Error("failed to get all delayed notifications in cron")
			return
		}

		notifier := NewNotifier(e.log, e.oneSignal)

		now := time.Now()
		for _, n := range allDelayedNotifications {
			timeDiff := now.Sub(n.TimeToNotify)
			if timeDiff.Minutes() <= 60 {
				err := notifier.Notify(n.Title, n.Comment, []string{n.UserDeviceID})
				if err != nil {
					e.log.WithError(err).Error("failed to notify users in cron")
				}
			}
		}
	})
	if err != nil {
		e.log.WithError(err).Error("failed to add new cron func into client")
		return
	}

	c.Start()
	defer c.Stop()
}
