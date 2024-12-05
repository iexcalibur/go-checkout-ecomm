package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iexcalibur/backend/internal/core/domain"
	"github.com/iexcalibur/backend/internal/core/ports"
)

type OrderService struct {
	orderRepo    ports.OrderRepository
	cartRepo     ports.CartRepository
	discountRepo ports.DiscountRepository
	discountSvc  *DiscountService
}

func NewOrderService(
	orderRepo ports.OrderRepository,
	cartRepo ports.CartRepository,
	discountRepo ports.DiscountRepository,
	discountSvc *DiscountService,
) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		cartRepo:     cartRepo,
		discountRepo: discountRepo,
		discountSvc:  discountSvc,
	}
}

func (s *OrderService) Checkout(userID string, discountCode string) (*domain.Order, *domain.Discount, error) {
	// 1. Get cart
	cart, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return nil, nil, err
	}

	if len(cart.Items) == 0 {
		return nil, nil, fmt.Errorf("cart is empty")
	}

	// Calculate total price for all items
	var totalAmount float64
	for _, item := range cart.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	// 2. Apply discount if code provided
	var discount *domain.Discount
	if discountCode != "" {
		discount, err = s.discountSvc.ValidateDiscountCode(discountCode)
		if err != nil {
			return nil, nil, err
		}
	}

	// 3. Create order
	order := &domain.Order{
		ID:          uuid.New().String(),
		UserID:      userID,
		Items:       cart.Items,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
	}

	// Apply discount to entire order if valid
	if discount != nil {
		order.DiscountID = discount.ID
		order.FinalAmount = totalAmount * (1 - discount.Percent/100)
		// Mark discount as used immediately
		if err := s.discountRepo.MarkAsUsed(discountCode); err != nil {
			return nil, nil, err
		}
	} else {
		order.FinalAmount = totalAmount
	}

	// 4. Save order
	if err := s.orderRepo.Create(*order); err != nil {
		return nil, nil, err
	}

	// 5. Clear cart
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, nil, err
	}

	// 6. Check if we should generate new discount code
	newDiscount, err := s.discountSvc.CheckAndGenerateDiscountForNthOrder()
	if err != nil {
		// Log error but don't fail the order
	}

	return order, newDiscount, nil
}
