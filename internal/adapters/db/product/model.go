package product

import (
	"database/sql"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres/types"
	"time"
)

type Model struct {
	Id              int                   `db:"id"`
	Title           string                `db:"title"`
	Description     sql.NullString        `db:"description"`
	Quantity        float32               `db:"quantity"`
	Thc             float32               `db:"thc"`
	Sativa          float32               `db:"sativa"`
	Indica          float32               `db:"indica"`
	ImagesUrl       types.NullStringArray `db:"images"`
	PricePerGram    float32               `db:"prices_per_gram"`
	SpecialDiscount float32               `db:"special_discount"`
	CreatedAt       time.Time             `db:"created_at"`
	UpdatedAt       sql.NullTime          `db:"updated_at"`
	DeletedAt       sql.NullTime          `db:"deleted_at"`
}

func (um Model) Map() product.Product {
	return product.Product{
		Id:              um.Id,
		Title:           um.Title,
		Description:     um.Description.String,
		Quantity:        um.Quantity,
		Thc:             um.Thc,
		Sativa:          um.Sativa,
		Indica:          um.Indica,
		ImagesUrl:       um.ImagesUrl.String,
		PricePerGram:    um.PricePerGram,
		SpecialDiscount: um.SpecialDiscount,
		CreatedAt:       um.CreatedAt,
		UpdatedAt:       um.UpdatedAt.Time,
		DeletedAt:       um.DeletedAt.Time,
	}
}
