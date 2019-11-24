package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/anodamobi/go-tg-api/bot"
)

type botResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type BotHandler struct {
	botSummary bot.Summary
	log        *logrus.Entry
}

func NewBotHandler(botSummary bot.Summary, log *logrus.Entry) *BotHandler {
	return &BotHandler{
		botSummary: botSummary,
		log:        log,
	}
}

func (h BotHandler) Handle(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(&botResponse{
		ID:       h.botSummary.ID,
		Name:     h.botSummary.Name,
		Username: h.botSummary.Username,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		h.log.WithError(err).Error("failed to serialize bot summary")
		return
	}

	_, _ = w.Write(response)
}
