package order

import "github.com/hramov/tg-bot-admin/internal/domain/product"

type InputOrderProduct struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type ShallowWeedProduct struct {
	InputOrderProduct
	PricesPerGram product.PriceForAmount `json:"prices_per_gram"`
	Title         string                 `json:"title"`
}

type ForMessage struct {
	Id         int                  `json:"id"`
	UserId     int                  `json:"user_info"`
	Products   []ShallowWeedProduct `json:"products"`
	TotalPrice int                  `json:"total_price"`
	Status     string               `json:"status"`
}

type InfoForUser struct {
	OrderId    int                  `json:"order_id"`
	Products   []ShallowWeedProduct `json:"products"`
	TotalPrice int                  `json:"total_price"`
	ChatId     int                  `json:"chat_id"`
}
