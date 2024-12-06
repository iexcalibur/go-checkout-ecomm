package ports

import "github.com/iexcalibur/backend/internal/models"

// CartRepository defines the interface for cart operations
type CartRepository interface {
	GetCart(userID string) (*models.Cart, error)
	AddToCart(userID string, item models.CartItem) error
}

// OrderRepository defines the interface for order operations
type OrderRepository interface {
	CreateOrder(order models.Order) error
	GetUserOrders(userID string) []models.Order
}

// PromoCodeRepository defines the interface for promo code operations
type PromoCodeRepository interface {
	CreatePromoCode(promoCode models.PromoCode) error
	GetPromoCode(code string) (*models.PromoCode, error)
}

// ProductRepository defines the interface for product operations
type ProductRepository interface {
	GetAll() []models.Product
	AddTestProduct(product models.Product)
}
