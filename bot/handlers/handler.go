package handlers

type BotHandler interface {
	LimitPriceSell(symbolId int64, price int64) (int, error)
	LimitPriceBuy(symbolId int64, price int64) (int, error)
	BestPriceSell(symbolId int64) (int, error)
	BestPriceBuy(symbolId int64) (int, error)
	GetCompanies() ([]Symbol, error)
	GetInfo() (Info, error)
}
