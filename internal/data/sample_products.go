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
			ImageURL:    "https://fansarmy.in/cdn/shop/products/anyafront_1800x1800.jpg?v=1659786556",
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
			ImageURL:    "https://www.snitch.co.in/cdn/shop/files/36dc00a9e8ffed42bdc6d27338f510b0.jpg?v=1730183553&width=1080",
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
			ImageURL:    "https://s3no.cashify.in/cashify/store/product//faecb80b21cd432d96bb82e7a59871ab.jpg?p=es5sq&s=es",
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
			ImageURL:    "https://veirdo.in/cdn/shop/files/imgpsh_fullsize_anim_31.jpg?v=1726139557&width=360",
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
			ImageURL:    "https://lp2.hm.com/hmgoepprod?set=quality%5B79%5D%2Csource%5B%2F23%2Fcf%2F23cff01e3ca2b8e5b4f040be0182a76bd4c0a3d2.jpg%5D%2Corigin%5Bdam%5D%2Ccategory%5Bladies_hoodiesswetshirts_hoodies%5D%2Ctype%5BDESCRIPTIVESTILLLIFE%5D%2Cres%5Bm%5D%2Chmver%5B2%5D&call=url[file:/product/main]",
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
			ImageURL:    "https://media-ik.croma.com/prod/https://media.croma.com/image/upload/v1694672652/Croma%20Assets/Entertainment/Wireless%20Earbuds/Images/301165_xzuxl0.png?tr=w-640",
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
			ImageURL:    "https://www.designinfo.in/wp-content/uploads/nc/p/5/2/4/1/7/52417-485x485-optimized.jpg'",
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
			ImageURL:    "https://store.storeimages.cdn-apple.com/1/as-images.apple.com/is/ipad-pro-finish-select-202405-11inch-silver-glossy-wifi?wid=5120&hei=2880&fmt=webp&qlt=70&.v=YXpaUEtKWGhlNnNrVGZkTEo4T0xsNEsrMGFueUl5dllOTm9xWTIwTHNieHBIK1RHU1d3MGRmSnYwdi9jcFlpODRqbGwveGJjbEtKKzd4aTZxV2lMQUFDb1F2RTNvUEVHRkpGaGtOSVFHalBTUFBNZ0M4MTl3aXZXQVpzNHkwVnkrWER2L1YvK1lCM2xaU21KNlhyaSt3&traceId=1",
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
			ImageURL:    "https://m.media-amazon.com/images/I/313vjNMPw3L._SX300_SY300_QL70_FMwebp_.jpg",
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
			ImageURL:    "https://m.media-amazon.com/images/I/41FVFskJChL._SX300_SY300_QL70_FMwebp_.jpg",
			Description: "Highly rated for durability and style.",
			Category:    "Others",
		},
	}

	for _, product := range products {
		store.AddTestProduct(product)
	}
}
