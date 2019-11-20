package user

import (
	"encoding/json"
	"net/http"
)

func (h *handler) GetCodes(w http.ResponseWriter, r *http.Request) {
	codes, _ := h.codeDB.GetAll()

	serializedBody, _ := json.Marshal(codes)

	_, _ = w.Write(serializedBody)
}
