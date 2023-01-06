package order

import (
	"github.com/hramov/tg-bot-admin/src/domain/product"
	"github.com/hramov/tg-bot-admin/src/domain/user"
	"time"
)

var registry Registry

type Status = string

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
	UserInfo   user.User            `json:"user_info"`
	Products   []ShallowWeedProduct `json:"products"`
	TotalPrice int                  `json:"total_price"`
	Status     Status               `json:"status"`
}

type InfoForUser struct {
	OrderId    int                  `json:"order_id"`
	Products   []ShallowWeedProduct `json:"products"`
	TotalPrice int                  `json:"total_price"`
	ChatId     int                  `json:"chat_id"`
}

type OutputOrder struct {
	Id         int                 `json:"id"`
	Status     Status              `json:"status"`
	TotalPrice int                 `json:"total_price"`
	UserInfo   user.User           `json:"user_info"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	SentAt     time.Time           `json:"sent_at"`
	DeletedAt  time.Time           `json:"deleted_at"`
	Products   []InputOrderProduct `json:"products"`
}
type Registry interface {
}

func New(r Registry) {
	if registry == nil {
		registry = r
	}
}
