package controller

import (
	"net/http"
	"strconv"
)

func (h *Handler) GetAccountInfo(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get(userHeader)

	idInt, err := strconv.Atoi(id)

	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.serv.User.Get(r.Context(), idInt)

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, info)
}

func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	info, err := h.serv.Company.GetAll(r.Context())

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, info)
}
