package errs

import (
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(
		[]byte(fmt.Sprintf(`{"code": %d, "error": "%s"}`, http.StatusBadRequest, err)),
	)
}

func InternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func Unauthorized(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write(
		[]byte(fmt.Sprintf(`{"code": %d, "error": "%s"}`, http.StatusUnauthorized, err)),
	)
}
