package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/anodamobi/go-tg-api/bot"
)

type homeResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type HomeHandler struct {
	botSummary bot.Summary
	log        *logrus.Entry
}

func NewHomeHandler(botSummary bot.Summary, log *logrus.Entry) *HomeHandler {
	return &HomeHandler{
		botSummary: botSummary,
		log:        log,
	}
}

func (h HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(&homeResponse{
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
