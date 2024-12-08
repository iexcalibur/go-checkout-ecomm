package domain

// CartService defines business logic for cart operations
type CartService interface {
	AddItem(userID string, item CartItem) error
	RemoveItem(userID string, productID string) error
	GetCart(userID string) (*Cart, error)
	Checkout(userID string, discountCode string) (*Order, error)
}

// CartRepository defines data access methods for cart
type CartRepository interface {
	Save(cart *Cart) error
	FindByUserID(userID string) (*Cart, error)
	Delete(userID string) error
}
