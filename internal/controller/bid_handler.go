package controller

import (
	"net/http"
	"strconv"
)

type PriceSellInput struct {
	CompanyId int `json:"company_id" validate:"required,gte=0"`
	Price     int `json:"price" validate:"required,gte=0"`
}

func (h *Handler) CreateSale(w http.ResponseWriter, r *http.Request) {

	var data PriceSellInput

	if ok := h.CheckRequest(w, r, &data); !ok {
		return
	}

	id := r.Header.Get(userHeader)

	idInt, err := strconv.Atoi(id)

	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.serv.Sale.Create(r.Context(), idInt, data.CompanyId, data.Price)

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, "created")
}

func (h *Handler) CreateBuy(w http.ResponseWriter, r *http.Request) {

	var data PriceSellInput

	if ok := h.CheckRequest(w, r, &data); !ok {
		return
	}

	id := r.Header.Get(userHeader)

	idInt, err := strconv.Atoi(id)

	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.serv.Buy.Create(r.Context(), idInt, data.CompanyId, data.Price)

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, "created")
}
