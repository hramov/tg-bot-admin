package db

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/filter"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"reflect"
	"strconv"
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
		if utils.Includes(tags, v.Field) == -1 && v.Field != "limit" && v.Field != "offset" {
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
	rawFilters := ctx.Value(filter.Key)
	if rawFilters == nil {
		return sql, nil, nil
	}
	filters := rawFilters.(filter.Options)
	var params []interface{}

	var paginationSql string
	var filterSql = " where "

	var limitStr string
	var offsetStr string

	for i, f := range filters {
		if f.Field == "limit" {
			limitStr = f.Value
			continue
		}
		if f.Field == "offset" {
			offsetStr = f.Value
			continue
		}
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

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return "", nil, err
		}
		paginationSql = fmt.Sprintf("limit $%d ", startN+1)
		params = append(params, limit)
	}

	if limitStr != "" && offsetStr != "" {
		_, err := strconv.Atoi(limitStr)
		if err != nil {
			return "", nil, err
		}

		offset, err := strconv.Atoi(limitStr)
		if err != nil {
			return "", nil, err
		}

		paginationSql += fmt.Sprintf("offset $%d", startN+2)
		params = append(params, offset)
	}

	return sql + filterSql + paginationSql, params, nil
}
