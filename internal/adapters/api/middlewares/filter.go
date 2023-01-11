package middlewares

import (
	"context"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/filter"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"net/http"
	"net/url"
	"strings"
)

func Filter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		unescape, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			utils.SendError(http.StatusBadRequest, err.Error(), w)
			return
		}
		queryArray := strings.Split(unescape, "&")
		if len(queryArray) == 0 || queryArray[0] == "" {
			h(w, r)
			return
		}
		var options filter.Options
		for _, v := range queryArray {
			option := filter.Option{}
			split := strings.Split(v, "=")
			option.Field = split[0]
			value := split[1]
			if strings.Contains(value, ":") {
				op := strings.Split(value, ":")[0]
				va := strings.Split(value, ":")[1]
				if utils.Includes(filter.Operators, op) != -1 {
					option.Operator = op
					option.Value = va
				} else {
					option.Operator = filter.Between
					option.Min = op
					option.Max = va
				}
			} else {
				option.Operator = filter.Eq
				option.Value = value
			}
			options = append(options, option)
		}
		ctx := context.WithValue(r.Context(), filter.Key, options)
		r = r.WithContext(ctx)
		h(w, r)
	}
}
