package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service"
	"io"
	"net/http"
)

type Handler struct {
	serv     *service.Services
	validate *validator.Validate
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{serv: services, validate: validator.New(validator.WithRequiredStructEnabled())}
}

func NewResponse(w http.ResponseWriter, status int, msg any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) CheckRequest(w http.ResponseWriter, r *http.Request, data any) bool {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, fmt.Sprintf("request body read error: %s", err.Error()))
		return false
	}

	err = json.Unmarshal(body, data)

	if err != nil {
		NewResponse(w, http.StatusBadRequest, fmt.Sprintf("unable to bind: %s", err.Error()))
		return false
	}

	err = h.validate.Struct(data)

	if err != nil {
		errs, ok := err.(validator.ValidationErrors)

		if ok {
			NewResponse(w, http.StatusBadRequest, errs.Error())
		} else {
			NewResponse(w, http.StatusBadRequest, err.Error())
		}

		return false
	}

	return true
}
