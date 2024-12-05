package memory

import (
	"sync"

	"github.com/iexcalibur/backend/internal/core/domain"
)

type MemoryCartRepository struct {
	carts map[string]domain.Cart
	mutex sync.RWMutex
}

func NewMemoryCartRepository() *MemoryCartRepository {
	return &MemoryCartRepository{
		carts: make(map[string]domain.Cart),
	}
}
