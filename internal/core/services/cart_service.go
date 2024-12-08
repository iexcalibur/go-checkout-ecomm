package services

import (
	"github.com/iexcalibur/backend/internal/core/domain"
)

type cartService struct {
	cartRepo  domain.CartRepository
	promoRepo domain.PromoRepository
}

func NewCartService(cartRepo domain.CartRepository, promoRepo domain.PromoRepository) domain.CartService {
	return &cartService{
		cartRepo:  cartRepo,
		promoRepo: promoRepo,
	}
}

func (s *cartService) AddItem(userID string, item domain.CartItem) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		cart = &domain.Cart{
			UserID: userID,
			Items:  []domain.CartItem{},
		}
	}

	// Add or update item
	found := false
	for i, existingItem := range cart.Items {
		if existingItem.ProductID == item.ProductID {
			cart.Items[i].Quantity += item.Quantity
			found = true
			break
		}
	}

	if !found {
		cart.Items = append(cart.Items, item)
	}

	// Recalculate total
	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	return s.cartRepo.Save(cart)
}

func (s *cartService) RemoveItem(userID string, productID string) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return err
	}

	var updatedItems []domain.CartItem
	for _, item := range cart.Items {
		if item.ProductID != productID {
			updatedItems = append(updatedItems, item)
		}
	}
	cart.Items = updatedItems

	// Recalculate total
	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	return s.cartRepo.Save(cart)
}

func (s *cartService) GetCart(userID string) (*domain.Cart, error) {
	return s.cartRepo.FindByUserID(userID)
}

func (s *cartService) Checkout(userID string, discountCode string) (*domain.Order, error) {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Create order and clear cart
	order := &domain.Order{
		UserID:      userID,
		TotalAmount: cart.Total,
		Items:       []domain.OrderItem{},
	}

	s.cartRepo.Delete(userID)
	return order, nil
}
