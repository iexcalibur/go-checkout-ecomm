package models

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
