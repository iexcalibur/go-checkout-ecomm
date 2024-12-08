package domain

// OrderService defines business logic for order operations
type OrderService interface {
	CreateOrder(cart *Cart, discountCode string) (*Order, error)
	GetUserOrders(userID string) ([]Order, error)
}

// OrderRepository defines data access methods for orders
type OrderRepository interface {
	Save(order *Order) error
	FindByUserID(userID string) ([]Order, error)
}
