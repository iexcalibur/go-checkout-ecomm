package storage

import (
	"testing"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore_AddToCart(t *testing.T) {
	store := NewMemoryStore()

	t.Run("Add new item to cart", func(t *testing.T) {
		item := models.CartItem{
			ItemID:   "1",
			Name:     "Test Product",
			Price:    100.0,
			Quantity: 2,
		}

		err := store.AddToCart("user1", item)
		assert.NoError(t, err)

		cart, err := store.GetCart("user1")
		assert.NoError(t, err)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, float64(200.0), cart.TotalPrice)
	})

	t.Run("Add existing item to cart", func(t *testing.T) {
		item := models.CartItem{
			ItemID:   "1",
			Name:     "Test Product",
			Price:    100.0,
			Quantity: 1,
		}

		err := store.AddToCart("user1", item)
		assert.NoError(t, err)

		cart, err := store.GetCart("user1")
		assert.NoError(t, err)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, 3, cart.Items[0].Quantity)
		assert.Equal(t, float64(300.0), cart.TotalPrice)
	})
}

func TestMemoryStore_RemoveFromCart(t *testing.T) {
	store := NewMemoryStore()

	// Setup test data
	item := models.CartItem{
		ItemID:   "1",
		Name:     "Test Product",
		Price:    100.0,
		Quantity: 2,
	}
	store.AddToCart("user1", item)

	t.Run("Remove existing item", func(t *testing.T) {
		err := store.RemoveFromCart("user1", "1")
		assert.NoError(t, err)

		cart, err := store.GetCart("user1")
		assert.Error(t, err)
		assert.Nil(t, cart)
	})

	t.Run("Remove from non-existent cart", func(t *testing.T) {
		err := store.RemoveFromCart("invalid_user", "1")
		assert.Error(t, err)
	})
}
