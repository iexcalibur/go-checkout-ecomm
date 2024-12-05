package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iexcalibur/backend/internal/storage"
)

type ProductHandler struct {
	store *storage.MemoryStore
}

func NewProductHandler(store *storage.MemoryStore) *ProductHandler {
	return &ProductHandler{store: store}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.store.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
