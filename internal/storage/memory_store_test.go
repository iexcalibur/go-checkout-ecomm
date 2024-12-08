package storage

import (
	"testing"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore_AddToCart(t *testing.T) {
	store := NewMemoryStore()

	item := models.CartItem{
		ProductID: "1",
		Name:      "Test Product",
		Price:     26.0,
		Quantity:  2,
	}

	err := store.AddToCart("user123", item)
	assert.NoError(t, err)

	cart, err := store.GetCart("user123")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cart.Items))
	assert.Equal(t, float64(52), cart.Total)
}

func TestMemoryStore_CreateOrder(t *testing.T) {
	store := NewMemoryStore()

	order := models.Order{
		UserID: "user123",
		Items: []models.OrderItem{
			{
				ProductID: "1",
				Name:      "Test Product",
				Price:     26.0,
				Quantity:  2,
			},
		},
		TotalAmount: 52.0,
	}

	err := store.CreateOrder(order)
	assert.NoError(t, err)

	orders := store.GetUserOrders("user123")
	assert.Equal(t, 1, len(orders))
}

func TestMemoryStore_PromoCodeOperations(t *testing.T) {
	store := NewMemoryStore()

	// Test creating promo code
	promo := models.PromoCode{
		Code:         "TEST10",
		DiscountRate: 10.0,
	}
	err := store.CreatePromoCode(promo)
	assert.NoError(t, err)

	// Test getting promo code
	savedPromo, err := store.GetPromoCode("TEST10")
	assert.NoError(t, err)
	assert.Equal(t, promo.Code, savedPromo.Code)

	// Test getting active promo code
	activePromo, err := store.GetActivePromoCode()
	assert.NoError(t, err)
	assert.Equal(t, promo.Code, activePromo.Code)

	// Test updating promo code
	savedPromo.Used = true
	err = store.UpdatePromoCode(*savedPromo)
	assert.NoError(t, err)

	// Test getting active promo code after using it
	_, err = store.GetActivePromoCode()
	assert.Error(t, err)
}

func TestMemoryStore_AutomaticPromoGeneration(t *testing.T) {
	store := NewMemoryStore()

	// Create 5 orders
	for i := 0; i < 5; i++ {
		order := models.Order{
			UserID:      "user123",
			TotalAmount: 100.0,
		}
		err := store.CreateOrder(order)
		assert.NoError(t, err)
	}

	// Check if automatic promo was generated
	promo, err := store.GetPromoCode("AUTO5")
	assert.NoError(t, err)
	assert.True(t, promo.IsAutomatic)
	assert.Equal(t, float64(10), promo.DiscountRate)
}
