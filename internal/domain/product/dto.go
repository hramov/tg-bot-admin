package product

import "time"

type InputProductImage = string
type OutputProductImage = string
type PriceForAmount map[string]int

type CoreProduct struct {
	Id              int                  `json:"id"`
	Title           string               `json:"title"`
	SpecialDiscount map[string]any       `json:"special_discount"`
	Quantity        int                  `json:"quantity"`
	Description     string               `json:"description"`
	Images          []OutputProductImage `json:"images"`
}

type InputWeedProduct struct {
	CoreProduct
	PricesPerGram PriceForAmount      `json:"prices_per_gram"`
	Indica        int                 `json:"indica"`
	Sativa        int                 `json:"sativa"`
	Thc           int                 `json:"thc"`
	Images        []InputProductImage `json:"images"`
}

type OutputCoreProduct struct {
	Id          int       `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type OutputWeedProduct struct {
	OutputCoreProduct
	PricesPerGram PriceForAmount `json:"title"`
	Thc           int            `json:"thc"`
	Indica        int            `json:"indica"`
	Sativa        int            `json:"sativa"`
}
