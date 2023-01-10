package api

import (
	"github.com/julienschmidt/httprouter"
	"time"
)

type ListenConfig struct {
	Type         string        `yaml:"type" env-default:"port"`
	BindIP       string        `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port         string        `yaml:"port" env-default:"8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	SockPath     string        `yaml:"sock_path" env-default:"app.sock"`
}

type CorsConfig struct {
	AllowedMethods     []string `yaml:"allowed_methods" env-required:"true"`
	AllowedOrigins     []string `yaml:"allowed_origins" env-required:"true"`
	AllowCredentials   bool     `yaml:"allow_credentials" env-required:"true"`
	AllowedHeaders     []string `yaml:"allowed_headers" env-required:"true"`
	OptionsPassthrough bool     `yaml:"options_passthrough" env-required:"true"`
	ExposedHeaders     []string `yaml:"exposed_headers" env-required:"true"`
	Debug              bool     `yaml:"debug" env-default:"false"`
}

type Handler interface {
	Init(router *httprouter.Router)
}
