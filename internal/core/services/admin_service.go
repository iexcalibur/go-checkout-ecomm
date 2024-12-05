package services

import (
	"github.com/iexcalibur/backend/internal/core/domain"
	"github.com/iexcalibur/backend/internal/core/ports"
)

type AdminService struct {
	orderRepo    ports.OrderRepository
	discountRepo ports.DiscountRepository
}

func NewAdminService(
	orderRepo ports.OrderRepository,
	discountRepo ports.DiscountRepository,
) *AdminService {
	return &AdminService{
		orderRepo:    orderRepo,
		discountRepo: discountRepo,
	}
}

// GetStats returns comprehensive statistics about orders and discounts
func (s *AdminService) GetStats() domain.OrderStats {
	return s.orderRepo.GetOrderStats()
}

// GetAllDiscounts returns all discount codes (used and unused)
func (s *AdminService) GetAllDiscounts() []domain.Discount {
	return s.discountRepo.GetAll()
}

// GenerateDiscountCode allows admin to manually generate a discount code
func (s *AdminService) GenerateDiscountCode() (*domain.Discount, error) {
	discountSvc := NewDiscountService(s.discountRepo, s.orderRepo, 10) // 10 is nth order
	return discountSvc.GenerateDiscountCode()
}

// GetActiveDiscountCode returns the current active discount code
func (s *AdminService) GetActiveDiscountCode() (*domain.Discount, error) {
	discountSvc := NewDiscountService(s.discountRepo, s.orderRepo, 10)
	return discountSvc.GetActiveDiscountCode()
}
