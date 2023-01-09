package config

import (
	"github.com/hramov/tg-bot-admin/pkg/db/postgres"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/hramov/tg-bot-admin/pkg/mail"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

const configPath = "config.yml"

type Config struct {
	IsDebug       *bool           `yaml:"is_debug" env-required:"true"`
	IsDevelopment *bool           `yaml:"is_development" env-required:"true"`
	Listen        ListenConfig    `yaml:"listen"`
	Cors          CorsConfig      `yaml:"cors"`
	Storage       postgres.Config `yaml:"storage"`
	Jwt           jwt.Config      `yaml:"jwt"`
	Mail          mail.Config     `yaml:"mail"`
	Logger        logging.Config  `yaml:"logger"`
}

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

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Println(err)
		}
	})
	return instance
}
