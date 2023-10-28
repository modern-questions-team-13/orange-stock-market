package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service"
	"net/http"
)

type Handler struct {
	serv     *service.Services
	validate *validator.Validate
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{serv: services, validate: validator.New()}
}

func NewResponse(w http.ResponseWriter, status int, msg any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
