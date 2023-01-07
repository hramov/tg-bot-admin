package product

import "time"

type Product struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Quantity        float32   `json:"quantity"`
	Thc             float32   `json:"thc"`
	Sativa          float32   `json:"sativa"`
	Indica          float32   `json:"indica"`
	ImagesUrl       []string  `json:"images"`
	PricePerGram    float32   `json:"prices_per_gram"`
	SpecialDiscount float32   `json:"special_discount"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
