package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
)

type CartHandler struct {
	store *storage.MemoryStore
}

func NewCartHandler(store *storage.MemoryStore) *CartHandler {
	return &CartHandler{store: store}
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	cart, err := h.store.GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var item struct {
		UserID    string  `json:"userId"`
		ProductID string  `json:"productId"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
		Name      string  `json:"name"`
		ImageURL  string  `json:"imageUrl"`
	}
	json.NewDecoder(r.Body).Decode(&item)

	cartItem := models.CartItem{
		ProductID: item.ProductID,
		Name:      item.Name,
		Price:     item.Price,
		Quantity:  item.Quantity,
		ImageURL:  item.ImageURL,
	}

	err := h.store.AddToCart(item.UserID, cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
	w.WriteHeader(http.StatusNotImplemented)
}
