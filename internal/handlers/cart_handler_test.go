package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestCartHandler_AddToCart(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewCartHandler(store)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "valid add to cart",
			payload: map[string]interface{}{
				"userId":    "user123",
				"productId": "1",
				"quantity":  2,
				"price":     26.0,
				"name":      "Test Product",
				"imageUrl":  "test.jpg",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid payload",
			payload:    map[string]interface{}{},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/cart/add", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.AddToCart(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestCartHandler_GetCart(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewCartHandler(store)

	tests := []struct {
		name       string
		userID     string
		wantStatus int
	}{
		{
			name:       "empty cart",
			userID:     "user123",
			wantStatus: http.StatusOK,
		},
		{
			name:       "missing user id",
			userID:     "",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/cart?userId="+tt.userID, nil)
			w := httptest.NewRecorder()

			handler.GetCart(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestCartHandler_Checkout(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewCartHandler(store)

	// Add item to cart first
	cartItem := models.CartItem{
		ProductID: "1",
		Name:      "Test Product",
		Price:     26.0,
		Quantity:  2,
	}
	store.AddToCart("user123", cartItem)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "valid checkout",
			payload: map[string]interface{}{
				"userId": "user123",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "checkout with promo",
			payload: map[string]interface{}{
				"userId":       "user123",
				"discountCode": "PROMO10",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid payload",
			payload:    map[string]interface{}{},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/cart/checkout", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.Checkout(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
