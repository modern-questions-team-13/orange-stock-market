package controller

import (
	"net/http"
	"strconv"
)

const tokenHeader = "Token"
const userHeader = "User"

type authMiddleware struct {
	handler *Handler
}

func newAuthMiddleware(handler *Handler) *authMiddleware {
	return &authMiddleware{handler: handler}
}

// Middleware function, which will be called for each request
func (m *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get(tokenHeader)

		id, ok := m.handler.serv.GetUserId(r.Context(), token)
		if !ok {
			NewResponse(w, http.StatusForbidden, "invalid token")
		} else {
			r.Header.Set(userHeader, strconv.Itoa(id))
			next.ServeHTTP(w, r)
		}
	})
}
