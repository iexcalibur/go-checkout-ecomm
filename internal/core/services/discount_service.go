package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iexcalibur/backend/internal/core/domain"
	"github.com/iexcalibur/backend/internal/core/ports"
)

type DiscountService struct {
	discountRepo ports.DiscountRepository
	orderRepo    ports.OrderRepository
	nthOrder     int // Configure which order number generates discount
	mutex        sync.Mutex
}

func NewDiscountService(
	discountRepo ports.DiscountRepository,
	orderRepo ports.OrderRepository,
	nthOrder int,
) *DiscountService {
	return &DiscountService{
		discountRepo: discountRepo,
		orderRepo:    orderRepo,
		nthOrder:     nthOrder,
	}
}

// Used by admin to generate discount code
func (s *DiscountService) GenerateDiscountCode() (*domain.Discount, error) {
	// Check if there's any unused discount code
	discounts := s.discountRepo.GetAll()
	for _, d := range discounts {
		if !d.Used {
			return nil, fmt.Errorf("there is already an active discount code available")
		}
	}

	discount := domain.Discount{
		ID:        uuid.New().String(),
		Code:      fmt.Sprintf("DISC%d", time.Now().Unix()),
		Percent:   10.0, // 10% discount
		Used:      false,
		CreatedAt: time.Now(),
	}

	if err := s.discountRepo.Create(discount); err != nil {
		return nil, err
	}

	return &discount, nil
}

// Used during checkout to validate discount code
func (s *DiscountService) ValidateDiscountCode(code string) (*domain.Discount, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	discount, err := s.discountRepo.GetByCode(code)
	if err != nil {
		return nil, fmt.Errorf("invalid discount code")
	}

	if discount.Used {
		return nil, fmt.Errorf("discount code already used")
	}

	// Ensure this is the latest available discount code
	allDiscounts := s.discountRepo.GetAll()
	var latestUnused *domain.Discount
	for _, d := range allDiscounts {
		if !d.Used && (latestUnused == nil || d.CreatedAt.After(latestUnused.CreatedAt)) {
			latestUnused = &d
		}
	}

	if latestUnused == nil || latestUnused.Code != code {
		return nil, fmt.Errorf("this discount code is no longer valid")
	}

	// Mark as used immediately to prevent race conditions
	if err := s.discountRepo.MarkAsUsed(code); err != nil {
		return nil, fmt.Errorf("failed to mark discount code as used")
	}

	return discount, nil
}

// Called after successful order to check if new discount should be generated
func (s *DiscountService) CheckAndGenerateDiscountForNthOrder() (*domain.Discount, error) {
	orderCount := s.orderRepo.GetOrderCount()

	// If it's the nth order, generate a new discount code
	if orderCount > 0 && orderCount%s.nthOrder == 0 {
		// Check if there's any unused discount code
		discounts := s.discountRepo.GetAll()
		for _, d := range discounts {
			if !d.Used {
				return nil, nil // Don't generate new code if there's an unused one
			}
		}
		return s.GenerateDiscountCode()
	}

	return nil, nil
}

// GetActiveDiscountCode returns the current unused discount code if available
func (s *DiscountService) GetActiveDiscountCode() (*domain.Discount, error) {
	discounts := s.discountRepo.GetAll()
	for _, d := range discounts {
		if !d.Used {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("no active discount code available")
}
