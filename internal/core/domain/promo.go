package domain

// PromoService defines business logic for promo code operations
type PromoService interface {
	GenerateCode(rate float64) (*PromoCode, error)
	ValidateCode(code string) (*PromoCode, error)
	GetActiveCode() (*PromoCode, error)
}

// PromoRepository defines data access methods for promo codes
type PromoRepository interface {
	Save(promo *PromoCode) error
	FindByCode(code string) (*PromoCode, error)
	FindActive() (*PromoCode, error)
}
