package auth

import (
	"github.com/anodamobi/glance-backend/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	userID  = "user_id"
)

func GetIDFromJWT(r *http.Request, auth *config.Authentication) (int64, string, error) {
	var tokenRaw string
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		tokenRaw = bearer[7:]
	}

	token, err := jwt.Parse(tokenRaw, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(auth.VerifyKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("cannot cast token.Claims to jwt.MapClaims")
	}

	rawUserID, ok := claims[userID]
	if !ok {
		return 0, "", errors.New("user_id is absent in the jwt")
	}

	userID, ok := rawUserID.(float64)
	if !ok {
		return 0, "", errors.New("failed to cast user_id into int")
	}

	return int64(userID), token.Raw, nil
}
