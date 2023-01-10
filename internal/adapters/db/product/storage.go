package product

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type productStorage struct {
	db     db.Connector
	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db db.Connector) product.IStorage {
	return &productStorage{db: db, logger: logger}
}

func (p *productStorage) GetBy(ctx context.Context, field string, param any) (*product.Product, error) {
	sql := fmt.Sprintf("select * from products where %s = $1", field)
	var params = []interface{}{param}
	res, err := db.ExecOne[product.Product, Model](ctx, p.db, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *productStorage) Get(ctx context.Context) ([]*product.Product, error) {
	sql := `select products.* from products`
	valid := db.ValidateFilters(ctx, &product.Product{})
	if !valid {
		return nil, fmt.Errorf("filters not valid")
	}
	sql, filterParams, err := db.FormatSqlFilters(sql, "products", 1, ctx)
	var params []interface{}
	for _, v := range filterParams {
		params = append(params, v)
	}
	res, err := db.Exec[product.Product, Model](ctx, p.db, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *productStorage) Create(ctx context.Context, dto product.InputWeedProduct) (*int, error) {
	sql := `
		insert into products 
			(title, description, quantity, thc, sativa, indica, images, prices_for_gram, special_discount, created_at, updated_at) 
		values 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), null)
		returning id`

	var params = []interface{}{dto.Title, dto.Description, dto.Quantity, dto.Thc, dto.Sativa, dto.Indica, dto.Images, dto.PricesPerGram, dto.SpecialDiscount}
	res, err := db.ExecOne[product.Product, Model](ctx, p.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (p *productStorage) Update(ctx context.Context, dto product.InputWeedProduct) (*int, error) {
	sql := `
		update users set
			title = $1,
			description = $2,
			quantity = $3,
			thc = $4,
			sativa = $5,
			indica = $6
			images = &7
			prices_for_gram = $8
			special_discount = $9
			updated_at = now()
		where id = $10
		returning id
	`
	var params = []interface{}{dto.Title, dto.Description, dto.Quantity, dto.Thc, dto.Sativa, dto.Indica, dto.Images, dto.PricesPerGram, dto.SpecialDiscount, dto.Id}
	res, err := db.ExecOne[product.Product, Model](ctx, p.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (p *productStorage) Delete(ctx context.Context, id int) (*int, error) {
	sql := `delete from products where id = $1 returning id`
	var params = []interface{}{id}
	res, err := db.ExecOne[product.Product, Model](ctx, p.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}
