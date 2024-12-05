package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
)

func TestCartHandler_AddToCart(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewCartHandler(store)

	// Add test product using the new method
	store.AddTestProduct(models.Product{
		ID:    "1",
		Name:  "Test Product",
		Price: 100.0,
	})

	tests := []struct {
		name       string
		request    CartRequest
		wantStatus int
	}{
		{
			name: "Valid add to cart",
			request: CartRequest{
				UserID:   "user1",
				ItemID:   "1",
				Quantity: 2,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Invalid item ID",
			request: CartRequest{
				UserID:   "user1",
				ItemID:   "999",
				Quantity: 1,
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req := httptest.NewRequest("POST", "/api/cart/add", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.AddToCart(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("AddToCart() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestCartHandler_Checkout(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewCartHandler(store)

	// Add test product using the new method
	store.AddTestProduct(models.Product{
		ID:    "1",
		Name:  "Test Product",
		Price: 100.0,
	})

	// Add item to cart
	cartReq := CartRequest{
		UserID:   "user1",
		ItemID:   "1",
		Quantity: 2,
	}
	cartItem := models.CartItem{
		ItemID:   cartReq.ItemID,
		Quantity: cartReq.Quantity,
	}
	store.AddToCart(cartReq.UserID, cartItem)

	tests := []struct {
		name       string
		request    CheckoutRequest
		wantStatus int
	}{
		{
			name: "Valid checkout without discount",
			request: CheckoutRequest{
				UserID: "user1",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Checkout with invalid user",
			request: CheckoutRequest{
				UserID: "invalid_user",
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req := httptest.NewRequest("POST", "/api/cart/checkout", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.Checkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Checkout() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}
