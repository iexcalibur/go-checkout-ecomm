package services

import (
	"errors"

	"github.com/iexcalibur/backend/internal/core/domain"
	"github.com/iexcalibur/backend/internal/core/ports"
)

type CartService struct {
	cartRepo     ports.CartRepository
	discountRepo ports.DiscountRepository
	orderRepo    ports.OrderRepository
}

func NewCartService(
	cartRepo ports.CartRepository,
	discountRepo ports.DiscountRepository,
	orderRepo ports.OrderRepository,
) *CartService {
	return &CartService{
		cartRepo:     cartRepo,
		discountRepo: discountRepo,
		orderRepo:    orderRepo,
	}
}

func (s *CartService) AddToCart(userID string, item domain.CartItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return s.cartRepo.AddItem(userID, item)
}

func (s *CartService) GetCart(userID string) (*domain.Cart, error) {
	return s.cartRepo.GetCart(userID)
}

func (s *CartService) RemoveFromCart(userID, itemID string) error {
	return s.cartRepo.RemoveItem(userID, itemID)
}

func (s *CartService) ApplyDiscount(userID, discountCode string) error {
	discount, err := s.discountRepo.GetByCode(discountCode)
	if err != nil {
		return err
	}

	if discount.Used {
		return errors.New("discount code already used")
	}

	cart, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return err
	}

	cart.DiscountID = discount.ID
	cart.FinalPrice = cart.TotalPrice * (1 - discount.Percent/100)

	return nil
}
