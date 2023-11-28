package kafka

type RequestMessage struct {
	Type      int    `json:"type"`
	CompanyId int    `json:"companyId"`
	Price     int    `json:"price"`
	DateTime  string `json:"datetime"`
}
