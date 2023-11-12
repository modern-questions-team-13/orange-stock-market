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

	middleware := newAuthMiddleware(h)

	noAuthR := router.PathPrefix("/signup").Subrouter()
	noAuthR.HandleFunc("", h.CreateUser).Methods(http.MethodPost)

	authR := router.PathPrefix("/api").Subrouter()
	authR.Use(middleware.Middleware)

	authR.HandleFunc("/LimitPriceBuy", h.CreateBuy).Methods(http.MethodPost)
	authR.HandleFunc("/LimitPriceSell", h.CreateSale).Methods(http.MethodPost)

	return &Router{Router: router}
}
