package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iexcalibur/backend/internal/models"
)

type Store struct {
	mu         sync.RWMutex
	products   map[string]models.Product
	carts      map[string]models.Cart
	orders     map[string]models.Order
	promoCodes map[string]models.PromoCode
}

func NewStore() *Store {
	return &Store{
		products:   make(map[string]models.Product),
		carts:      make(map[string]models.Cart),
		orders:     make(map[string]models.Order),
		promoCodes: make(map[string]models.PromoCode),
	}
}

// Product methods
func (s *Store) AddProduct(product models.Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products[product.ID] = product
}

func (s *Store) GetAllProducts() []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	products := make([]models.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products
}

// Cart methods
func (s *Store) GetCart(userID string) (*models.Cart, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, cart := range s.carts {
		if cart.UserID == userID {
			return &cart, nil
		}
	}
	return nil, fmt.Errorf("cart not found for user: %s", userID)
}

func (s *Store) AddToCart(userID string, item models.CartItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var cart models.Cart
	found := false

	for _, c := range s.carts {
		if c.UserID == userID {
			cart = c
			found = true
			break
		}
	}

	if !found {
		cart = models.Cart{
			ID:     uuid.New().String(),
			UserID: userID,
			Items:  []models.CartItem{},
		}
	}

	itemFound := false
	for i, existingItem := range cart.Items {
		if existingItem.ProductID == item.ProductID {
			cart.Items[i].Quantity += item.Quantity
			itemFound = true
			break
		}
	}

	if !itemFound {
		cart.Items = append(cart.Items, item)
	}

	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	s.carts[cart.ID] = cart
	return nil
}

// Order methods
func (s *Store) CreateOrder(order models.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	s.orders[order.ID] = order
	return nil
}

func (s *Store) GetUserOrders(userID string) []models.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var userOrders []models.Order
	for _, order := range s.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}

// PromoCode methods
func (s *Store) CreatePromoCode(promoCode models.PromoCode) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	promoCode.ID = uuid.New().String()
	promoCode.GeneratedAt = time.Now()
	s.promoCodes[promoCode.Code] = promoCode
	return nil
}

func (s *Store) GetPromoCode(code string) (*models.PromoCode, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if promo, exists := s.promoCodes[code]; exists {
		return &promo, nil
	}
	return nil, fmt.Errorf("promo code not found: %s", code)
}
