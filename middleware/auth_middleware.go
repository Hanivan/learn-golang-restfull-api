package middleware

import (
	"Hanivan/learn-golang-restfull-api/helper"
	"Hanivan/learn-golang-restfull-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if apiKey := r.Header.Get("X-API-Key"); apiKey == "RAHASIA" {
		m.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZE",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
