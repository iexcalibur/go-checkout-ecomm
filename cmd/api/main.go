package main

import (
	"log"

	"github.com/iexcalibur/backend/internal/config"
	"github.com/iexcalibur/backend/internal/core/middleware"
	"github.com/iexcalibur/backend/internal/data"
	"github.com/iexcalibur/backend/internal/handlers"
	"github.com/iexcalibur/backend/internal/routes"
	"github.com/iexcalibur/backend/internal/storage"
)

func main() {
	// Initialize server
	server := config.NewServer("8080")

	// Initialize storage and sample data
	store := storage.NewMemoryStore()
	data.InitializeSampleProducts(store)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(store)
	cartHandler := handlers.NewCartHandler(store)
	orderHandler := handlers.NewOrderHandler(store)
	promoHandler := handlers.NewPromoHandler(store)

	// Setup routes
	routes.SetupRoutes(
		server.Router(),
		productHandler,
		cartHandler,
		orderHandler,
		promoHandler,
	)

	// Apply middleware
	router := server.Router()
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
