package order

import (
	"github.com/hramov/tg-bot-admin/internal/domain/order"
)

type Model struct {
}

func (um Model) Map() order.Order {
	return order.Order{}
}
