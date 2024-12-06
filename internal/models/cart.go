package models

type Cart struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Items     []CartItem `json:"items"`
	Total     float64    `json:"total"`
	Discount  float64    `json:"discount"`
	PromoCode *PromoCode `json:"promo_code,omitempty"`
}

type CartItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ImageURL  string  `json:"image_url"`
}
