package handlers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type CtxKey int

const (
	logCtxKey = iota
	httpCtxKey
	dbCtxKey
	jwtCtxKey
)

func CtxLog(entry *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(logCtxKey).(*logrus.Entry)
}

func CtxHTTP(http *url.URL) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, httpCtxKey, http)
	}
}

func HTTP(r *http.Request) *url.URL {
	return r.Context().Value(httpCtxKey).(*url.URL)
}

func CtxJWT(entry *jwtauth.JWTAuth) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtCtxKey, entry)
	}
}

func JWT(r *http.Request) *jwtauth.JWTAuth {
	return r.Context().Value(jwtCtxKey).(*jwtauth.JWTAuth)
}
