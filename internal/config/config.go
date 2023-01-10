package config

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/hramov/tg-bot-admin/pkg/mail"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

const configPath = "config.yml"

type Config struct {
	IsDebug       *bool            `yaml:"is_debug" env-required:"true"`
	IsDevelopment *bool            `yaml:"is_development" env-required:"true"`
	Listen        api.ListenConfig `yaml:"listen"`
	Cors          api.CorsConfig   `yaml:"cors"`
	Storage       postgres.Config  `yaml:"storage"`
	Jwt           jwt.Config       `yaml:"jwt"`
	Mail          mail.Config      `yaml:"mail"`
	Logger        logging.Config   `yaml:"logger"`
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
