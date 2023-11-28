package handlers

import "io"

type BotHandler interface {
	LimitPriceSell(symbolId int64, price int64) (int, io.ReadCloser, error)
	LimitPriceBuy(symbolId int64, price int64) (int, io.ReadCloser, error)
	BestPriceSell(symbolId int64) (int, error)
	BestPriceBuy(symbolId int64) (int, error)
	GetCompanies() ([]Symbol, error)
	GetInfo() (Info, error)
}
