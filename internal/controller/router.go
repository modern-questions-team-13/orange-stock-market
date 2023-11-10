package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func NewRouter(h *Handler) *Router {
	router := mux.NewRouter()

	//middleware := newAuthMiddleware(h)

	noAuthR := router.PathPrefix("/signup").Subrouter()
	noAuthR.HandleFunc("", h.CreateUser).Methods(http.MethodPost)

	//authR.Use(middleware.Middleware)

	return &Router{Router: router}
}
