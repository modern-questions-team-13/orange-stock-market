package model

import "time"

type Buy struct {
	Id        int       `json:"id"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int       `json:"user_id"`
	CompanyId int       `json:"company_id"`
}
