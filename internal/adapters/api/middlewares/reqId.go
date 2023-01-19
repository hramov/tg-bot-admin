package middlewares

import (
	"context"
	"github.com/google/uuid"
	"github.com/hramov/tg-bot-admin/internal/config"
	"net/http"
)

func ReqId(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		w.Header().Set(config.RequestId, id.String())
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), config.RequestId, id)))
	})
}
