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
	promo := models.PromoCode{
		Code:         "PROMO10",
		DiscountRate: 10.0,
	}
	err := h.store.CreatePromoCode(promo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(promo)
}

func (h *PromoHandler) GetActivePromoCode(w http.ResponseWriter, r *http.Request) {
	promo, err := h.store.GetPromoCode("PROMO10")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(promo)
}
