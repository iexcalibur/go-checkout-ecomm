package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
)

type PromoHandler struct {
	store *storage.MemoryStore
}

func NewPromoHandler(store *storage.MemoryStore) *PromoHandler {
	return &PromoHandler{store: store}
}

func (h *PromoHandler) GeneratePromoCode(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Code         string  `json:"code"`
		DiscountRate float64 `json:"discount_rate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if request.Code == "" {
		http.Error(w, "Promo code cannot be empty", http.StatusBadRequest)
		return
	}
	if request.DiscountRate <= 0 || request.DiscountRate > 100 {
		http.Error(w, "Discount rate must be between 0 and 100", http.StatusBadRequest)
		return
	}

	promo := models.PromoCode{
		Code:         request.Code,
		DiscountRate: request.DiscountRate,
	}

	err := h.store.CreatePromoCode(promo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promo)
}

func (h *PromoHandler) GetActivePromoCode(w http.ResponseWriter, r *http.Request) {
	promo, err := h.store.GetActivePromoCode()
	if err != nil {
		http.Error(w, "No active promo code found", http.StatusNotFound)
		return
	}

	response := struct {
		*models.PromoCode
		Message string `json:"message,omitempty"`
	}{
		PromoCode: promo,
	}

	if promo.IsAutomatic {
		response.Message = "Congratulations! You've earned this promo code after 5 orders!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
