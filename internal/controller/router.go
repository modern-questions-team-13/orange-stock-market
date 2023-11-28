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

	compR := router.PathPrefix("/companies").Subrouter()
	compR.HandleFunc("", h.GetCompanies).Methods(http.MethodGet)

	authR := router.PathPrefix("/api").Subrouter()
	authR.Use(middleware.Middleware)

	authR.HandleFunc("/LimitPriceBuy", h.CreateBuy).Methods(http.MethodPost)
	authR.HandleFunc("/LimitPriceSell", h.CreateSale).Methods(http.MethodPost)

	authR.HandleFunc("/info", h.GetAccountInfo).Methods(http.MethodGet)
	authR.HandleFunc("/getCompanies", h.GetCompanies).Methods(http.MethodGet)

	return &Router{Router: router}
}
