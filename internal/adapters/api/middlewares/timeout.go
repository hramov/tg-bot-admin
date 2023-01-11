package middlewares

import (
	"net/http"
)

// Timeout TODO fix panic when deadline exceeded
func Timeout(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ctx, _ := context.WithTimeout(r.Context(), config.DefaultTimeout)
		//defer cancel()
		//r = r.WithContext(ctx)
		h(w, r)
	}
}
