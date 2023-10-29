package model

type Operation struct {
	Id        int `json:"id"`
	Price     int `json:"price"`
	CreatedAt int `json:"created_at"`
	BuyerId   int `json:"user_id"`
	SellerId  int `json:"seller_id"`
	CompanyId int `json:"company_id"`
}
