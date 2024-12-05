package storage

import (
	"errors"
	"sync"

	"github.com/iexcalibur/backend/internal/models"
)

type MemoryStore struct {
	products  map[string]models.Product
	carts     map[string]models.Cart
	discounts map[string]models.Discount
	mutex     sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		products:  make(map[string]models.Product),
		carts:     make(map[string]models.Cart),
		discounts: make(map[string]models.Discount),
	}
}

func (m *MemoryStore) GetCart(userID string) (*models.Cart, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	cart, exists := m.carts[userID]
	if !exists {
		return nil, errors.New("cart not found")
	}
	return &cart, nil
}

func (m *MemoryStore) ClearCart(userID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.carts, userID)
	return nil
}

func (m *MemoryStore) ValidateDiscount(code string) (*models.Discount, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, discount := range m.discounts {
		if discount.Code == code && !discount.Used {
			return &discount, nil
		}
	}
	return nil, errors.New("invalid or used discount code")
}

func (m *MemoryStore) AddToCart(userID string, item models.CartItem) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	cart, exists := m.carts[userID]
	if !exists {
		cart = models.Cart{
			UserID: userID,
			Items:  make([]models.CartItem, 0),
		}
	}

	// Check if item already exists in cart
	for i, existingItem := range cart.Items {
		if existingItem.ItemID == item.ItemID {
			cart.Items[i].Quantity += item.Quantity
			cart.TotalPrice += item.Price * float64(item.Quantity)
			m.carts[userID] = cart
			return nil
		}
	}

	// Add new item
	cart.Items = append(cart.Items, item)
	cart.TotalPrice += item.Price * float64(item.Quantity)
	m.carts[userID] = cart
	return nil
}

func (m *MemoryStore) RemoveFromCart(userID string, itemID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	cart, exists := m.carts[userID]
	if !exists {
		return errors.New("cart not found")
	}

	for i, item := range cart.Items {
		if item.ItemID == itemID {
			cart.TotalPrice -= item.Price * float64(item.Quantity)
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			m.carts[userID] = cart
			return nil
		}
	}

	return errors.New("item not found in cart")
}

func (m *MemoryStore) UpdateCartItemQuantity(userID string, itemID string, quantity int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	cart, exists := m.carts[userID]
	if !exists {
		return errors.New("cart not found")
	}

	for i, item := range cart.Items {
		if item.ItemID == itemID {
			cart.TotalPrice -= item.Price * float64(item.Quantity)
			cart.Items[i].Quantity = quantity
			cart.TotalPrice += item.Price * float64(quantity)
			m.carts[userID] = cart
			return nil
		}
	}

	return errors.New("item not found in cart")
}

// Product methods
func (m *MemoryStore) GetAll() []models.Product {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	products := make([]models.Product, 0, len(m.products))
	for _, product := range m.products {
		products = append(products, product)
	}
	return products
}

func (m *MemoryStore) GetByID(id string) (models.Product, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	product, exists := m.products[id]
	if !exists {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (m *MemoryStore) Update(product models.Product) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.products[product.ID]; !exists {
		return errors.New("product not found")
	}

	m.products[product.ID] = product
	return nil
}

func (m *MemoryStore) Delete(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(m.products, id)
	return nil
}

// AddTestProduct is used only for testing
func (m *MemoryStore) AddTestProduct(product models.Product) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.products[product.ID] = product
}
