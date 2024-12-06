package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iexcalibur/backend/internal/storage"
)

type OrderHandler struct {
	store *storage.MemoryStore
}

func NewOrderHandler(store *storage.MemoryStore) *OrderHandler {
	return &OrderHandler{store: store}
}

func (h *OrderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	orders := h.store.GetUserOrders(userID)
	json.NewEncoder(w).Encode(orders)
}
