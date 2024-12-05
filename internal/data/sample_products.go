package data

import (
	"github.com/iexcalibur/backend/internal/models"
	"github.com/iexcalibur/backend/internal/storage"
)

func InitializeSampleProducts(store *storage.MemoryStore) {
	products := []models.Product{
		{
			ID:          "1",
			Name:        "Spy x Family Tshirt",
			Price:       26,
			Rating:      4.6,
			Sold:        1015,
			Location:    "South Jakarta",
			ImageURL:    "https://via.placeholder.com/250",
			Description: "Highly rated for durability and style.",
			Category:    "Clothing",
		},
		{
			ID:          "2",
			Name:        "Green Man Jacket",
			Price:       49,
			Rating:      4.3,
			Sold:        1578,
			Location:    "Yogyakarta",
			ImageURL:    "https://via.placeholder.com/200",
			Description: "Best value for money product.",
			Category:    "Clothing",
		},
		{
			ID:          "3",
			Name:        "Iphone 14 Pro Max",
			Price:       1200,
			Rating:      4.3,
			Sold:        1532,
			Location:    "Yogyakarta",
			ImageURL:    "https://via.placeholder.com/350",
			Description: "Best value for money product.",
			Category:    "Electronics",
		},
		{
			ID:          "4",
			Name:        "Oversized Tshirt",
			Price:       48,
			Rating:      4.1,
			Sold:        1406,
			Location:    "North Purwokerto",
			ImageURL:    "https://via.placeholder.com/250",
			Description: "Best value for money product.",
			Category:    "Clothing",
		},
		{
			ID:          "5",
			Name:        "Brown Woman Hoodie",
			Price:       49,
			Rating:      4.3,
			Sold:        710,
			Location:    "South Jakarta",
			ImageURL:    "https://via.placeholder.com/150",
			Description: "A must-have item for enthusiasts.",
			Category:    "Clothing",
		},
		{
			ID:          "6",
			Name:        "Airpod Pro 2022",
			Price:       459,
			Rating:      3.6,
			Sold:        1841,
			Location:    "South Jakarta",
			ImageURL:    "https://via.placeholder.com/350",
			Description: "Top-quality product with excellent features.",
			Category:    "Electronics",
		},
		{
			ID:          "7",
			Name:        "DJI Mini 3 Pro",
			Price:       842,
			Rating:      3.8,
			Sold:        923,
			Location:    "South Jakarta",
			ImageURL:    "https://via.placeholder.com/300",
			Description: "Customer favorite for its design and functionality.",
			Category:    "Electronics",
		},
		{
			ID:          "8",
			Name:        "Ipad Pro Gen 3",
			Price:       338,
			Rating:      4.1,
			Sold:        1803,
			Location:    "South Jakarta",
			ImageURL:    "https://via.placeholder.com/300",
			Description: "Top-quality product with excellent features.",
			Category:    "Electronics",
		},
		{
			ID:          "9",
			Name:        "G502 X Lightspeed Wireless Gaming Mouse",
			Price:       139,
			Rating:      4.3,
			Sold:        697,
			Location:    "North Purwokerto",
			ImageURL:    "https://via.placeholder.com/300",
			Description: "Highly rated for durability and style.",
			Category:    "Computer Accessories",
		},
		{
			ID:          "10",
			Name:        "Logitech Refurbished G920",
			Price:       1280,
			Rating:      4.6,
			Sold:        1293,
			Location:    "Yogyakarta",
			ImageURL:    "https://via.placeholder.com/200",
			Description: "Highly rated for durability and style.",
			Category:    "Others",
		},
	}

	for _, product := range products {
		store.AddTestProduct(product)
	}
}
