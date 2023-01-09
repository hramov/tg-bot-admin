package product

import (
	"github.com/hramov/tg-bot-admin/internal/domain/product"
)

type Model struct {
}

func (um Model) Map() product.Product {
	return product.Product{}
}
