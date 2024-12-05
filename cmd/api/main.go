package main

import (
	"log"

	"github.com/iexcalibur/backend/internal/config"
	"github.com/iexcalibur/backend/internal/core/services"
	"github.com/iexcalibur/backend/internal/data"
	"github.com/iexcalibur/backend/internal/handlers"
	"github.com/iexcalibur/backend/internal/repositories/memory"
	"github.com/iexcalibur/backend/internal/routes"
	"github.com/iexcalibur/backend/internal/storage"
)

func main() {
	// Initialize server
	server := config.NewServer("8080")

	// Initialize repositories
	orderRepo := memory.NewMemoryOrderRepository()
	discountRepo := memory.NewMemoryDiscountRepository()

	// Initialize storage and sample data
	store := storage.NewMemoryStore()
	data.InitializeSampleProducts(store)

	// Initialize services
	adminService := services.NewAdminService(orderRepo, discountRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(store)
	cartHandler := handlers.NewCartHandler(store)
	adminHandler := handlers.NewAdminHandler(adminService)

	// Setup routes
	routes.SetupRoutes(
		server.Router(),
		productHandler,
		cartHandler,
		adminHandler,
	)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
