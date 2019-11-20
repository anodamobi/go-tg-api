package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

// Authenticator is a middleware's method which check and handle errors from validate function
func Authenticator(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return authenticateRequest(jwtSecret, TokenFromCookie, TokenFromQuery, TokenFromHeader)(next)
	}
}

func authenticateRequest(jwtSecret string, findTokenFns ...func(r *http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			token, err := authenticate(jwtSecret, r, findTokenFns...)
			if err != nil {
				http.SetCookie(w, &http.Cookie{
					Name:  "jwt",
					Value: "",
					Path:  "/",
				})

				http.Error(w, errors.Errorf("failed to authenticate token, the reason is: %s", err.Error()).Error(), http.StatusUnauthorized)

				return
			}

			ctx := context.WithValue(r.Context(), jwtauth.TokenCtxKey, token)

			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Value: token.Raw,
				Path:  "/",
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(handlerFunc)
	}
}

func authenticate(jwtSecret string, r *http.Request, findTokenFns ...func(r *http.Request) string) (*jwt.Token, error) {
	var tokenStr string
	for _, fn := range findTokenFns {
		tokenStr = fn(r)
		if tokenStr != "" {
			break
		}
	}
	if tokenStr == "" {
		return nil, errors.New("cannot find token")
	}

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	var ok bool
	token.Claims, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast token.Claims to jwt.MapClaims")
	}

	token.Raw, err = token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenFromCookie tries to retreive the token string from a cookie named
// "jwt".
func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// TokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER T".
func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

// TokenFromQuery tries to retreive the token string from the "jwt" URI
// query parameter.
func TokenFromQuery(r *http.Request) string {
	// Get token from query param named "jwt".
	return r.URL.Query().Get("jwt")
}
