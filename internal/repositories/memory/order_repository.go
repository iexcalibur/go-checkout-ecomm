package memory

import (
	"fmt"
	"sync"

	"github.com/iexcalibur/backend/internal/core/domain"
)

type MemoryOrderRepository struct {
	orders map[string]domain.Order // key: order ID
	mutex  sync.RWMutex
	stats  domain.OrderStats
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[string]domain.Order),
		stats: domain.OrderStats{
			TotalOrders:    0,
			TotalAmount:    0,
			DiscountCodes:  make([]string, 0),
			TotalDiscounts: 0,
			ItemsPurchased: 0,
		},
	}
}

func (r *MemoryOrderRepository) Create(order domain.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.orders[order.ID] = order

	// Update statistics
	r.stats.TotalOrders++
	r.stats.TotalAmount += order.TotalAmount

	// Calculate discount amount and update stats
	if order.DiscountID != "" {
		discountAmount := order.TotalAmount - order.FinalAmount
		r.stats.TotalDiscounts += discountAmount
		r.stats.DiscountCodes = append(r.stats.DiscountCodes, order.DiscountID)
	}

	// Count total items purchased
	for _, item := range order.Items {
		r.stats.ItemsPurchased += item.Quantity
	}

	return nil
}

func (r *MemoryOrderRepository) GetOrderCount() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.stats.TotalOrders
}

func (r *MemoryOrderRepository) GetOrderStats() domain.OrderStats {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.stats
}

// Additional helper methods for getting order history
func (r *MemoryOrderRepository) GetOrdersByUserID(userID string) []domain.Order {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	userOrders := make([]domain.Order, 0)
	for _, order := range r.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}

func (r *MemoryOrderRepository) GetOrderByID(orderID string) (domain.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if order, exists := r.orders[orderID]; exists {
		return order, nil
	}
	return domain.Order{}, fmt.Errorf("order not found")
}
