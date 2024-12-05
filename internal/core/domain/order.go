package domain

import "time"

type Order struct {
	ID          string     `json:"id"`
	UserID      string     `json:"userId"`
	Items       []CartItem `json:"items"`
	TotalAmount float64    `json:"totalAmount"`
	DiscountID  string     `json:"discountId,omitempty"`
	FinalAmount float64    `json:"finalAmount"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type OrderStats struct {
	TotalOrders    int      `json:"totalOrders"`
	TotalAmount    float64  `json:"totalAmount"`
	DiscountCodes  []string `json:"discountCodes"`
	TotalDiscounts float64  `json:"totalDiscounts"`
	ItemsPurchased int      `json:"itemsPurchased"`
}
