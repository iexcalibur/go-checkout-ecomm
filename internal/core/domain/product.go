package domain

// ProductService defines business logic for product operations
type ProductService interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id string) (*Product, error)
}

// ProductRepository defines data access methods for products
type ProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(id string) (*Product, error)
}
