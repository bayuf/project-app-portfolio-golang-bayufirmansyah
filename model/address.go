package model

type Address struct {
	Model      Model
	UserId     int    `json:"user_id"`
	Detail     string `json:"detail"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
