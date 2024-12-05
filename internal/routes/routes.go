package routes

import (
	"github.com/gorilla/mux"
	"github.com/iexcalibur/backend/internal/handlers"
)

func SetupRoutes(
	r *mux.Router,
	productHandler *handlers.ProductHandler,
	cartHandler *handlers.CartHandler,
	adminHandler *handlers.AdminHandler,
) {
	// Product route
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET")

	// Cart routes
	r.HandleFunc("/api/cart/add", cartHandler.AddToCart).Methods("POST")
	r.HandleFunc("/api/cart", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/api/cart/{itemId}", cartHandler.RemoveFromCart).Methods("DELETE")
	r.HandleFunc("/api/cart/{itemId}/quantity", cartHandler.UpdateCartItemQuantity).Methods("PUT")
	r.HandleFunc("/api/cart/checkout", cartHandler.Checkout).Methods("POST")

	// Admin routes
	r.HandleFunc("/api/admin/discount/active", adminHandler.GetActiveDiscountCode).Methods("GET")
	r.HandleFunc("/api/admin/stats", adminHandler.GetStats).Methods("GET")
	r.HandleFunc("/api/admin/discount", adminHandler.GenerateDiscountCode).Methods("POST")
}
