package db

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/filter"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"reflect"
	"strings"
)

func Exec[T any, V Mapper[T]](ctx context.Context, db Connector, sql string, params []interface{}) ([]*T, error) {
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)
	var model []V
	err = conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}
	var dto []*T
	for _, v := range model {
		d := v.Map()
		dto = append(dto, &d)
	}
	return dto, nil
}

func ExecOne[T any, V Mapper[T]](ctx context.Context, db Connector, sql string, params []interface{}) (*T, error) {
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)
	var model []V
	err = conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}
	if len(model) == 0 {
		return new(T), nil
	}
	dto := model[0].Map()
	return &dto, nil
}

func ValidateFilters[T any](ctx context.Context, model *T) bool {
	f := ctx.Value(filter.Key)
	if f == nil {
		return true
	}
	filters := f.(filter.Options)
	var tags []string
	elem := reflect.TypeOf(model).Elem()
	n := elem.NumField()
	for i := 0; i < n; i++ {
		field := elem.Field(i)
		tag := field.Tag.Get("json")
		jTag := strings.Split(tag, ",")[0]
		tags = append(tags, jTag)
	}
	for _, v := range filters {
		if utils.Includes(tags, v.Field) == -1 {
			return false
		}
	}
	return true
}

func formatFilterOperator(f filter.Option, tName string, n int) (string, []interface{}, int, error) {
	switch f.Operator {
	case filter.Eq:
		return fmt.Sprintf("%s.%s = $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.NotEq:
		return fmt.Sprintf("%s.%s <> $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.LowerThan:
		return fmt.Sprintf("%s.%s < $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.LowerThanEqual:
		return fmt.Sprintf("%s.%s <= $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.GreaterThan:
		return fmt.Sprintf("%s.%s > $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.GreaterThanEqual:
		return fmt.Sprintf("%s.%s >= $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	case filter.Between:
		return fmt.Sprintf("%s.%s >= $%d and %s.%s < $%d", tName, f.Field, n, tName, f.Field, n+1), []interface{}{f.Min, f.Max}, n + 1, nil
	case filter.Like:
		return fmt.Sprintf("%s.%s ilike $%d", tName, f.Field, n), []interface{}{f.Value}, n, nil
	default:
		return "", nil, n, fmt.Errorf("no such operator")
	}
}

func FormatSqlFilters(sql string, tName string, startN int, ctx context.Context) (string, []interface{}, error) {
	f := ctx.Value(filter.Key)
	if f == nil {
		return sql, nil, nil
	}
	filters := f.(filter.Options)
	var params []interface{}
	filterSql := " where "
	for i, f := range filters {
		s, p, n, err := formatFilterOperator(f, tName, i+startN)
		startN = n
		if err != nil {
			return "", nil, err
		}
		filterSql += fmt.Sprintf("%s and ", s)
		for _, v := range p {
			params = append(params, v)
		}
	}
	filterSql = utils.CutStrSuffix(filterSql, 4)
	return sql + filterSql, params, nil
}
