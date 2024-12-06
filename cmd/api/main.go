package main

import (
	"log"

	"github.com/iexcalibur/backend/internal/config"
	"github.com/iexcalibur/backend/internal/data"
	"github.com/iexcalibur/backend/internal/handlers"
	"github.com/iexcalibur/backend/internal/routes"
	"github.com/iexcalibur/backend/internal/storage"
	"github.com/rs/cors"
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

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
	})

	// Use CORS handler in server
	server.SetHandler(c.Handler(server.Router()))

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
