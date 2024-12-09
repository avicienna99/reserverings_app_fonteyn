package db

type House struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	Availability string `json:"availability"`
}
