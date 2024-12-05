package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
)

type CartHandler struct {
	store *storage.MemoryStore
}

func NewCartHandler(store *storage.MemoryStore) *CartHandler {
	return &CartHandler{store: store}
}

type CartRequest struct {
	UserID   string `json:"userId"`
	ItemID   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type CheckoutRequest struct {
	UserID       string `json:"userId"`
	DiscountCode string `json:"discountCode,omitempty"`
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req CartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cartItem := models.CartItem{
		ItemID:   req.ItemID,
		Quantity: req.Quantity,
	}

	if err := h.store.AddToCart(req.UserID, cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cart, err := h.store.GetCart(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	cart, err := h.store.GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["itemId"]
	userID := r.URL.Query().Get("userId")

	if userID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	if err := h.store.RemoveFromCart(userID, itemID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) UpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["itemId"]

	var req CartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.store.UpdateCartItemQuantity(req.UserID, itemID, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	cart, err := h.store.GetCart(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cart, err := h.store.GetCart(req.UserID)
	if err != nil {
		http.Error(w, "Failed to get cart", http.StatusInternalServerError)
		return
	}

	order := models.Order{
		ID:          uuid.New().String(),
		UserID:      req.UserID,
		Items:       cart.Items,
		TotalAmount: cart.TotalPrice,
		CreatedAt:   time.Now(),
	}

	if req.DiscountCode != "" {
		discount, err := h.store.ValidateDiscount(req.DiscountCode)
		if err != nil {
			http.Error(w, "Invalid discount code", http.StatusBadRequest)
			return
		}
		order.DiscountID = discount.ID
		order.FinalAmount = order.TotalAmount * 0.9
	} else {
		order.FinalAmount = order.TotalAmount
	}

	if err := h.store.ClearCart(req.UserID); err != nil {
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
