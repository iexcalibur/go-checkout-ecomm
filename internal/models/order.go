package models

import "time"

type Order struct {
	ID             string      `json:"id"`
	UserID         string      `json:"user_id"`
	Items          []OrderItem `json:"items"`
	TotalAmount    float64     `json:"total_amount"`
	DiscountCode   string      `json:"discount_code"`
	DiscountAmount float64     `json:"discount_amount"`
	CreatedAt      time.Time   `json:"created_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
