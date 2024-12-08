package handlers

import (
	"encoding/json"
	"net/http"

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

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	cart, err := h.store.GetCart(userID)
	if err != nil {
		// If cart doesn't exist, return an empty cart
		cart = &models.Cart{
			UserID: userID,
			Items:  []models.CartItem{},
		}
	}

	w.Header().Set("Content-Type", "application/json")
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

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
	vars := mux.Vars(r)
	productID := vars["productId"]

	var update struct {
		UserID   string `json:"userId"`
		Quantity int    `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cart, err := h.store.GetCart(update.UserID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	found := false
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity = update.Quantity
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Product not found in cart", http.StatusNotFound)
		return
	}

	// Recalculate total
	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	// Update cart in store
	if err := h.store.UpdateCart(*cart); err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	userID := r.URL.Query().Get("userId")

	cart, err := h.store.GetCart(userID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	// Find and remove the item
	found := false
	var updatedItems []models.CartItem
	for _, item := range cart.Items {
		if item.ProductID != productID {
			updatedItems = append(updatedItems, item)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Product not found in cart", http.StatusNotFound)
		return
	}

	cart.Items = updatedItems

	// Recalculate total
	cart.Total = 0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}

	// Update cart in store
	if err := h.store.UpdateCart(*cart); err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var checkout struct {
		UserID       string `json:"userId"`
		DiscountCode string `json:"discountCode,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the user's cart
	cart, err := h.store.GetCart(checkout.UserID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	if len(cart.Items) == 0 {
		http.Error(w, "Cart is empty", http.StatusBadRequest)
		return
	}

	// Apply discount if provided
	var discountAmount float64
	if checkout.DiscountCode != "" {
		promoCode, err := h.store.GetPromoCode(checkout.DiscountCode)
		if err != nil {
			http.Error(w, "Invalid discount code", http.StatusBadRequest)
			return
		}
		if promoCode.Used {
			http.Error(w, "Discount code already used", http.StatusBadRequest)
			return
		}
		discountAmount = cart.Total * (promoCode.DiscountRate / 100)

		// Mark promo code as used
		promoCode.Used = true
		if err := h.store.UpdatePromoCode(*promoCode); err != nil {
			http.Error(w, "Failed to update promo code", http.StatusInternalServerError)
			return
		}
	}

	// Create order items from cart items
	var orderItems []models.OrderItem
	for _, item := range cart.Items {
		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Create the order
	order := models.Order{
		UserID:         checkout.UserID,
		Items:          orderItems,
		TotalAmount:    cart.Total,
		DiscountCode:   checkout.DiscountCode,
		DiscountAmount: discountAmount,
	}

	if err := h.store.CreateOrder(order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Clear the cart
	if err := h.store.ClearCart(checkout.UserID); err != nil {
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}

	// Return the created order
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
