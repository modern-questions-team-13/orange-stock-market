package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bids struct {
	Price    int64 `json:"Price"`
	Quantity int64 `json:"Quantity"`
}

type Stock struct {
	Id     int64  `json:"id"`
	Ticker string `json:"ticker"`
	Bids   []Bids `json:"bids"`
}

type Symbol struct {
	Id     int64  `json:"id"`
	Ticker string `json:"ticker"`
}

type LimitPriceBid struct {
	Id         int64   `json:"id"`
	Account    Account `json:"account"`
	SymbolId   int64   `json:"SymbolId"`
	Price      int64   `json:"Price"`
	Type       string  `json:"type"`
	CreateDate string  `json:"createDate"`
}

type Account struct {
	Id     int    `json:"id"`
	Login  string `json:"login"`
	Wealth int    `json:"wealth"`
}

type Asset struct {
	Id       int64 `json:"id"`
	Quantity int64 `json:"Quantity"`
}

type Info struct {
	Account Account `json:"account"`
	Assets  []Asset `json:"assets"`
}

func GetSellStock() ([]Stock, error) {
	req, err := http.NewRequest("GET", "https://datsorange.devteam.games/sellStock", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error Code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res []Stock
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	//log.Println(res, err)
	return res, nil
}

func (b *botHandler) GetCompanies() ([]Symbol, error) {
	req, err := http.NewRequest("GET", b.url+"/getCompanies", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Token", b.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error Code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res []Symbol
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetBuyStock() ([]Stock, error) {
	req, err := http.NewRequest("GET", "https://datsorange.devteam.games/buyStock", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR STATUCE %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res []Stock
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	//log.Println(res, err)
	return res, err
}

func (b *botHandler) GetInfo() (Info, error) {
	req, err := http.NewRequest("GET", b.url+"/info", nil)
	if err != nil {
		return Info{}, err
	}
	req.Header.Add("Token", b.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Info{}, err
	}
	if resp.StatusCode != 200 {
		return Info{}, fmt.Errorf("Code Error %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Info{}, err
	}
	var res Info
	err = json.Unmarshal(body, &res)
	if err != nil {
		return Info{}, err
	}
	//log.Println(res, err)
	return res, err
}
