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

func formatSqlFilters[T any](ctx context.Context, model *T, sql string, tName string, startN int) (string, []interface{}, error) {
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

	var sortBy string
	var descStr string

	for i, f := range filters {
		if f.Field == "count" {
			limitStr = f.Value
			continue
		}
		if f.Field == "start" {
			offsetStr = f.Value
			continue
		}

		if f.Field == "sortBy" {
			sortBy = f.Value
			continue
		}

		if f.Field == "desc" {
			descStr = f.Value
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

	if sortBy != "" && validateSortBy(model, sortBy) {
		if descStr != "" {
			if descStr == "true" {
				paginationSql += fmt.Sprintf("order by %s desc ", sortBy)
			} else {
				paginationSql += fmt.Sprintf("order by %s ", sortBy)
			}
		} else {
			paginationSql += fmt.Sprintf("order by %s ", sortBy)
		}
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return "", nil, err
		}
		paginationSql += fmt.Sprintf("limit $%d ", startN+1)
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

func validateFilters[T any](ctx context.Context, model *T) bool {
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
		tag := field.Tag.Get(filter.Tag)
		jTag := strings.Split(tag, ",")[0]
		tags = append(tags, jTag)
	}
	for _, v := range filters {
		if utils.Includes(tags, v.Field) == -1 && utils.Includes(filter.GeneralFilters, v.Field) == -1 {
			return false
		}
	}
	return true
}

func validateSortBy[T any](model *T, field string) bool {
	elem := reflect.TypeOf(model).Elem()
	n := elem.NumField()
	for i := 0; i < n; i++ {
		f := elem.Field(i)
		tag := f.Tag.Get(filter.Tag)
		jTag := strings.Split(tag, ",")[0]
		if field == jTag {
			return true
		}
	}
	return false
}
