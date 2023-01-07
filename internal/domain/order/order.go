package order

import (
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"time"
)

type Order struct {
	Id         int               `json:"id"`
	UserInfo   user.User         `json:"user_info"`
	Products   []product.Product `json:"products"`
	TotalPrice float32           `json:"total_price"`
	Status     string            `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  time.Time         `json:"deleted_at"`
	SentAt     time.Time         `json:"sent_at"`
	DoneAt     time.Time         `json:"done_at"`
}
