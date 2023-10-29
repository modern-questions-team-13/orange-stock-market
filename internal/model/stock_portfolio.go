package model

type StockPortfolio struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	CompanyId int `json:"company_id"`
	Count     int `json:"count"`
}
