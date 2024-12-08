package services

import (
	"time"

	"github.com/iexcalibur/backend/internal/core/domain"
)

type orderService struct {
	orderRepo domain.OrderRepository
	promoRepo domain.PromoRepository
}

func NewOrderService(orderRepo domain.OrderRepository, promoRepo domain.PromoRepository) domain.OrderService {
	return &orderService{
		orderRepo: orderRepo,
		promoRepo: promoRepo,
	}
}

func (s *orderService) CreateOrder(cart *domain.Cart, discountCode string) (*domain.Order, error) {
	order := &domain.Order{
		UserID:      cart.UserID,
		TotalAmount: cart.Total,
		CreatedAt:   time.Now(),
	}

	// Convert cart items to order items
	for _, item := range cart.Items {
		orderItem := domain.OrderItem{
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}

	// Apply discount if provided
	if discountCode != "" {
		promo, err := s.promoRepo.FindByCode(discountCode)
		if err == nil && !promo.Used {
			order.DiscountCode = discountCode
			order.DiscountAmount = order.TotalAmount * (promo.DiscountRate / 100)
			order.TotalAmount -= order.DiscountAmount
		}
	}

	return order, s.orderRepo.Save(order)
}

func (s *orderService) GetUserOrders(userID string) ([]domain.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}
