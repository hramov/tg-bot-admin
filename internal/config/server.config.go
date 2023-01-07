package config

import (
	"github.com/gin-contrib/cors"
	"time"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS", "DELETE"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"Content-Length", "X-Forwarded-For"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}
