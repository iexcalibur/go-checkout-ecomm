package domain

import "time"

type Discount struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Percent   float64   `json:"percent"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"createdAt"`
	UsedAt    time.Time `json:"usedAt,omitempty"`
}
