package model

type OperationType int

const (
	SellType OperationType = iota
	BuyType
)

type Operation struct {
	Id            int `json:"id"`
	Price         int `json:"price"`
	CreatedAt     int `json:"created_at"`
	UserId        int `json:"user_id"`
	CompanyId     int `json:"company_id"`
	OperationType int `json:"operation_type"`
}
