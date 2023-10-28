package controller

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter(h *Handler) *Router {
	router := mux.NewRouter()
	return &Router{Router: router}
}
