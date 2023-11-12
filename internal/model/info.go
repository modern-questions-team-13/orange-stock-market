package model

type Info struct {
	Account User    `json:"account"`
	Assets  []Asset `json:"assets"`
}
