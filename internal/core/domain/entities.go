package domain

import "time"

type Cart struct {
	ID        string
	UserID    string
	Items     []CartItem
	Total     float64
	Discount  float64
	PromoCode *PromoCode
}

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
	ImageURL  string
}

type Order struct {
	ID             string
	UserID         string
	Items          []OrderItem
	TotalAmount    float64
	DiscountCode   string
	DiscountAmount float64
	CreatedAt      time.Time
}

type OrderItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Product struct {
	ID          string
	Name        string
	Price       float64
	Rating      float64
	Sold        int
	Location    string
	ImageURL    string
	Description string
	Category    string
}

type PromoCode struct {
	ID           string
	Code         string
	DiscountRate float64
	Used         bool
	GeneratedAt  time.Time
	IsAutomatic  bool
	Expiry       *time.Time
}
