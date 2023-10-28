package model

type StockPortfolio struct {
	UserId    int `json:"user_id"`
	CompanyId int `json:"company_id"`
	Count     int `json:"count"`
}
