package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type botHandler struct {
	url   string
	token string
}

type cancelReq struct {
	bidId int64 `json:"bidId"`
}

type reqLimitPrice struct {
	SymbolId int64 `json:"company_id"`
	Price    int64 `json:"price"`
}

func NewBotHandler(url, token string) *botHandler {
	return &botHandler{
		url:   url,
		token: token,
	}
}

func (b *botHandler) Cancel(bidId int64) (int, error) {
	reqBody, err := json.Marshal(cancelReq{bidId: bidId})
	if err != nil {
		return 0, err
	}
	req, _ := http.NewRequest("POST", "https://datsorange.devteam.games/RemoveBid", bytes.NewBuffer(reqBody))
	req.Header.Add("token", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)

	return resp.StatusCode, nil

}

func (b *botHandler) LimitPriceSell(symbolId int64, price int64) (int, error) {
	reqBody, err := json.Marshal(reqLimitPrice{
		SymbolId: symbolId,
		Price:    price,
	})
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", b.url+"/LimitPriceSell", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Token", b.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, err
}

func (b *botHandler) LimitPriceBuy(symbolId int64, price int64) (int, error) {
	reqBody, err := json.Marshal(reqLimitPrice{
		SymbolId: symbolId,
		Price:    price,
	})
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", b.url+"/LimitPriceBuy", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Token", b.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil

}

func (b *botHandler) BestPriceSell(symbolId int64) (int, error) {
	reqBody, err := json.Marshal(reqLimitPrice{
		SymbolId: symbolId,
		Price:    0,
	})
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", b.url+"/LimitPriceSell", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Token", b.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil

}

func (b *botHandler) BestPriceBuy(symbolId int64) (int, error) {
	reqBody, err := json.Marshal(reqLimitPrice{
		SymbolId: symbolId,
		Price:    1e9,
	})
	if err != nil {
		return 0, err
	}
	req, _ := http.NewRequest("POST", b.token+"/LimitPriceBuy", bytes.NewBuffer(reqBody))
	req.Header.Add("Token", b.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)

	return resp.StatusCode, nil

}
