package controller

import (
	"github.com/modern-questions-team-13/orange-stock-market/internal/infrastructure/kafka"
	"net/http"
	"strconv"
	"time"
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

	err = h.serv.KafkaSenderService.SendMessage(kafka.RequestMessage{
		Type:      1,
		CompanyId: data.CompanyId,
		Price:     data.Price,
		DateTime:  time.Now().Format(time.RFC3339),
	})

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
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

	err = h.serv.KafkaSenderService.SendMessage(kafka.RequestMessage{
		Type:      0,
		CompanyId: data.CompanyId,
		Price:     data.Price,
		DateTime:  time.Now().Format(time.RFC3339),
	})

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.serv.Buy.Create(r.Context(), idInt, data.CompanyId, data.Price)

	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(w, http.StatusOK, "created")
}
