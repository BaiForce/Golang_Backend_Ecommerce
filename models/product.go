package models

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CategoryID  int     `json:"category_id"`
	Price       float64 `json:"price"`
	Description string  `json:"description"` // Tambahkan field description
}
