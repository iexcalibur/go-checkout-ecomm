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

func TestPromoHandler_GeneratePromoCode(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewPromoHandler(store)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name: "valid promo code",
			payload: map[string]interface{}{
				"code":          "SUMMER25",
				"discount_rate": 25.0,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "invalid discount rate",
			payload: map[string]interface{}{
				"code":          "TEST",
				"discount_rate": 150.0,
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/admin/discount", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.GeneratePromoCode(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestPromoHandler_GetActivePromoCode(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewPromoHandler(store)

	t.Run("no active promo", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/admin/discount/active", nil)
		w := httptest.NewRecorder()

		handler.GetActivePromoCode(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	// Create a promo code and test again
	store.CreatePromoCode(models.PromoCode{
		Code:         "TEST10",
		DiscountRate: 10.0,
	})

	t.Run("active promo exists", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/admin/discount/active", nil)
		w := httptest.NewRecorder()

		handler.GetActivePromoCode(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
