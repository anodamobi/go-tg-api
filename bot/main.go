package bot

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"

	"github.com/go-chi/jwtauth"

	"github.com/anodamobi/go-tg-api/config"

	"github.com/sirupsen/logrus"

	"github.com/anodamobi/go-tg-api/db"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

type Summary struct {
	ID       int
	Name     string
	Username string
}

type Boss struct {
	*tgbotapi.BotAPI
	db  *db.DB
	log *logrus.Entry
	jwt *jwtauth.JWTAuth
}

func NewBoss(bot *config.Bot, db *db.DB, log *logrus.Entry, jwt *jwtauth.JWTAuth) (*Boss, error) {
	botAPI, err := tgbotapi.NewBotAPI(bot.Token)
	if err != nil {
		return nil, err
	}

	return &Boss{
		BotAPI: botAPI,
		db:     db,
		log:    log,
		jwt:    jwt,
	}, nil
}

func (b Boss) Summary() Summary {
	var name = b.Self.FirstName
	if b.Self.LastName != "" {
		name = fmt.Sprintf("%s %s", name, b.Self.LastName)
	}

	return Summary{
		ID:       b.Self.ID,
		Name:     name,
		Username: b.Self.UserName,
	}
}

func (b *Boss) Listen() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.GetUpdatesChan(u)
	if err != nil {
		return errors.Wrap(err, "failed to get updates from Telegram")
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		user := update.Message.From
		switch update.Message.Command() {
		case startCmd:
			//store information about user
			photoID, err := b.getUserPhoto(user.ID)
			if err != nil {
				b.log.WithError(err).
					WithField("user_id", user.ID).
					Error("failed to get user photos")
				continue
			}

			//TODO: check if user already exists
			uid := uuid.NewV4()
			err = b.uploadUserInfo(user, photoID, uid.String())
			if err != nil {
				b.log.WithError(err).
					WithField("user_id", user.ID).
					Error("failed to upload user")
				continue
			}

			f, _ := b.BotAPI.GetFile(tgbotapi.FileConfig{
				FileID: photoID,
			})

			fmt.Println(f.Link(b.Token))

			_, token, _ := b.jwt.Encode(jwt.MapClaims{
				"id":  uid.String(),
				"exp": time.Now().Add(tokenExpirationDuration).Unix(),
			})
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, token)
			msg.ReplyToMessageID = update.Message.MessageID

			_, _ = b.Send(msg)
		}
	}

	return nil
}

func (b Boss) uploadUserInfo(user *tgbotapi.User, photoID string, uid string) error {
	return b.db.CreateUser(&db.User{
		ID:         uid,
		ExternalID: user.ID,
		Name:       fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Language:   user.LanguageCode,
		Avatar:     photoID,
		UpdatedAt:  time.Now(),
	})
}

func (b Boss) getUserPhoto(userID int) (string, error) {
	photos, err := b.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{
		UserID: userID,
		Offset: 1,
		Limit:  1,
	})

	if err != nil {
		return "", err
	}

	var photoID string
	if photos.TotalCount != 0 {
		photoID = photos.Photos[0][0].FileID
	}

	return photoID, nil
}
