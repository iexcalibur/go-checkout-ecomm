package routes

import (
	"github.com/gorilla/mux"
	"github.com/iexcalibur/backend/internal/handlers"
)

func SetupRoutes(
	router *mux.Router,
	productHandler *handlers.ProductHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
	promoHandler *handlers.PromoHandler,
) {
	// Product routes
	router.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET", "OPTIONS")

	// Cart routes
	router.HandleFunc("/api/cart", cartHandler.GetCart).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cart/add", cartHandler.AddToCart).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cart/{productId}", cartHandler.UpdateCartItem).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/cart/{productId}", cartHandler.RemoveFromCart).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/cart/checkout", cartHandler.Checkout).Methods("POST", "OPTIONS")

	// Order routes
	router.HandleFunc("/api/orders", orderHandler.GetUserOrders).Methods("GET", "OPTIONS")

	// Promo code routes
	router.HandleFunc("/api/admin/discount", promoHandler.GeneratePromoCode).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/admin/discount/active", promoHandler.GetActivePromoCode).Methods("GET", "OPTIONS")
}
