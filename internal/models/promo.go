package models

import "time"

type PromoCode struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	Code         string     `json:"code"`
	DiscountRate float64    `json:"discount_rate"`
	Used         bool       `json:"used"`
	GeneratedAt  time.Time  `json:"generated_at"`
	Expiry       *time.Time `json:"expiry,omitempty"`
}

type StoreSummary struct {
	TotalItemsPurchased int      `json:"total_items_purchased"`
	TotalRevenue        float64  `json:"total_revenue"`
	DiscountCodesUsed   []string `json:"discount_codes_used"`
	TotalDiscountAmount float64  `json:"total_discount_amount"`
}
