package user

import (
	"encoding/json"
	"net/http"
)

type Respnnnn struct {
	Users []User `json:"users"`
}

type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func (h *handler) GetTempPass(w http.ResponseWriter, r *http.Request) {
	users, _ := h.db.GetTempPass()

	var usersR []User
	for _, u := range users {
		usersR = append(usersR, User{
			Email: u.Email,
			Pass:  u.TempPass,
		})
	}

	serializedBody, _ := json.Marshal(Respnnnn{Users: usersR})

	_, _ = w.Write(serializedBody)
}
