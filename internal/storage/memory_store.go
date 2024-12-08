package storage

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iexcalibur/backend/internal/models"
)

type MemoryStore struct {
	mu         sync.RWMutex
	products   map[string]models.Product
	carts      map[string]models.Cart
	orders     map[string]models.Order
	promoCodes map[string]models.PromoCode
	orderCount int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		products:   make(map[string]models.Product),
		carts:      make(map[string]models.Cart),
		orders:     make(map[string]models.Order),
		promoCodes: make(map[string]models.PromoCode),
		orderCount: 0,
	}
}

func init() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
}

// Product methods
func (s *MemoryStore) AddTestProduct(product models.Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products[product.ID] = product
}

func (s *MemoryStore) GetAll() []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	products := make([]models.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products
}

// Cart methods
func (s *MemoryStore) GetCart(userID string) (*models.Cart, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, cart := range s.carts {
		if cart.UserID == userID {
			return &cart, nil
		}
	}
	return nil, fmt.Errorf("cart not found for user: %s", userID)
}

func (s *MemoryStore) AddToCart(userID string, item models.CartItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var cart models.Cart
	found := false

	// Find existing cart for user
	for _, c := range s.carts {
		if c.UserID == userID {
			cart = c
			found = true
			break
		}
	}

	// Create new cart if not found
	if !found {
		cart = models.Cart{
			ID:     uuid.New().String(),
			UserID: userID,
			Items:  []models.CartItem{},
		}
	}

	// Update or add item
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

	// Update total with rounding
	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += roundToTwoDecimals(item.Price * float64(item.Quantity))
	}
	cart.Total = roundToTwoDecimals(cart.Total)

	s.carts[cart.ID] = cart
	return nil
}

// Order methods
func (s *MemoryStore) CreateOrder(order models.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.TotalAmount = roundToTwoDecimals(order.TotalAmount)
	order.DiscountAmount = roundToTwoDecimals(order.DiscountAmount)
	s.orders[order.ID] = order
	s.orderCount++

	// Generate promo code if needed
	if err := s.generateAutomaticPromoCode(); err != nil {
		return err
	}

	return nil
}

func (s *MemoryStore) GetUserOrders(userID string) []models.Order {
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
func (s *MemoryStore) CreatePromoCode(promoCode models.PromoCode) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	promoCode.ID = uuid.New().String()
	promoCode.GeneratedAt = time.Now()
	s.promoCodes[promoCode.Code] = promoCode
	return nil
}

func (s *MemoryStore) GetPromoCode(code string) (*models.PromoCode, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if promo, exists := s.promoCodes[code]; exists {
		return &promo, nil
	}
	return nil, fmt.Errorf("promo code not found: %s", code)
}

func (s *MemoryStore) UpdateCart(cart models.Cart) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.carts[cart.ID]; !exists {
		return fmt.Errorf("cart not found")
	}

	s.carts[cart.ID] = cart
	return nil
}

func (s *MemoryStore) UpdatePromoCode(promoCode models.PromoCode) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.promoCodes[promoCode.Code]; !exists {
		return fmt.Errorf("promo code not found")
	}

	s.promoCodes[promoCode.Code] = promoCode
	return nil
}

func (s *MemoryStore) ClearCart(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var cartID string
	found := false

	// Find the cart for the user
	for id, cart := range s.carts {
		if cart.UserID == userID {
			cartID = id
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("cart not found for user")
	}

	// Delete the cart
	delete(s.carts, cartID)
	return nil
}

func (s *MemoryStore) GetActivePromoCode() (*models.PromoCode, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Find the most recent unused promo code
	var activePromo *models.PromoCode
	var automaticPromo *models.PromoCode

	for _, promo := range s.promoCodes {
		if promo.Used {
			continue
		}

		// Handle automatic promos separately
		if promo.IsAutomatic {
			if automaticPromo == nil || promo.GeneratedAt.After(automaticPromo.GeneratedAt) {
				promoTemp := promo
				automaticPromo = &promoTemp
			}
			continue
		}

		// Handle manual promos
		if activePromo == nil || promo.GeneratedAt.After(activePromo.GeneratedAt) {
			promoTemp := promo
			activePromo = &promoTemp
		}
	}

	// Prefer automatic promos over manual ones
	if automaticPromo != nil {
		return automaticPromo, nil
	}

	if activePromo != nil {
		return activePromo, nil
	}

	return nil, fmt.Errorf("no active promo code found")
}

func (s *MemoryStore) generateAutomaticPromoCode() error {
	// Check if we should generate a promo code
	if s.orderCount%3 != 0 {
		return nil
	}

	// Generate random code
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := "PROMO"

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	code += string(b)

	// Create and save promo code
	promo := models.PromoCode{
		ID:           uuid.New().String(),
		Code:         code,
		DiscountRate: 10.0,
		GeneratedAt:  time.Now(),
		IsAutomatic:  true,
	}

	// Store directly since we already have the lock from CreateOrder
	s.promoCodes[promo.Code] = promo
	return nil
}

func roundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
