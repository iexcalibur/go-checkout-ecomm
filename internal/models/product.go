package models

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	Sold        int     `json:"sold"`
	Location    string  `json:"location"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}
