package model

type Secret struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
