package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/iexcalibur/backend/internal/core/domain"
)

type MemoryDiscountRepository struct {
	discounts map[string]domain.Discount // key: discount code
	mutex     sync.RWMutex
}

func NewMemoryDiscountRepository() *MemoryDiscountRepository {
	return &MemoryDiscountRepository{
		discounts: make(map[string]domain.Discount),
	}
}

func (r *MemoryDiscountRepository) Create(discount domain.Discount) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.discounts[discount.Code]; exists {
		return errors.New("discount code already exists")
	}

	r.discounts[discount.Code] = discount
	return nil
}

func (r *MemoryDiscountRepository) GetByCode(code string) (*domain.Discount, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if discount, exists := r.discounts[code]; exists {
		return &discount, nil
	}
	return nil, errors.New("discount code not found")
}

func (r *MemoryDiscountRepository) MarkAsUsed(code string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if discount, exists := r.discounts[code]; exists {
		discount.Used = true
		discount.UsedAt = time.Now()
		r.discounts[code] = discount
		return nil
	}
	return errors.New("discount code not found")
}

func (r *MemoryDiscountRepository) GetAll() []domain.Discount {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	discounts := make([]domain.Discount, 0, len(r.discounts))
	for _, discount := range r.discounts {
		discounts = append(discounts, discount)
	}
	return discounts
}
