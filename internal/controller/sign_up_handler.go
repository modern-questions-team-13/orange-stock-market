package controller

import (
	"errors"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
	"net/http"
)

type UserCreateInputData struct {
	Login  string `json:"login" validate:"required,min=1,max=255"`
	Wealth int    `json:"wealth" validate:"required,gte=0"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var data UserCreateInputData

	if ok := h.CheckRequest(w, r, &data); !ok {
		return
	}

	token, err := h.serv.User.Create(r.Context(), data.Login, data.Wealth)

	if errors.Is(err, repoerrs.ErrAlreadyExists) {
		NewResponse(w, http.StatusConflict, err.Error())
		return
	}

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, token)
}
