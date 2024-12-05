package ports

import "github.com/iexcalibur/backend/internal/core/domain"

type CartRepository interface {
	AddItem(userID string, item domain.CartItem) error
	GetCart(userID string) (*domain.Cart, error)
	RemoveItem(userID, itemID string) error
	ClearCart(userID string) error
	UpdateItemQuantity(userID, itemID string, quantity int) error
}

type OrderRepository interface {
	Create(order domain.Order) error
	GetOrderCount() int
	GetOrderStats() domain.OrderStats
}

type DiscountRepository interface {
	Create(discount domain.Discount) error
	GetByCode(code string) (*domain.Discount, error)
	MarkAsUsed(code string) error
	GetAll() []domain.Discount
}
