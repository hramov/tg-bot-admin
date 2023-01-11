package middlewares

import (
	"net/http"
)

func Error(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
