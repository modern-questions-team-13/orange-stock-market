package model

type BidInfo struct {
	UserId    int `json:"user_id"`
	CompanyId int `json:"company_id"`
	Price     int `json:"price"`
}
