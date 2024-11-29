package db

// House represents a house entity in the database
type House struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}
