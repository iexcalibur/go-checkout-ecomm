package models

type Cart struct {
	UserID     string     `json:"userId"`
	Items      []CartItem `json:"items"`
	TotalPrice float64    `json:"totalPrice"`
}

type CartItem struct {
	ItemID   string  `json:"itemId"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
