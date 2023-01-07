package config

import (
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

const configPath = "config.yml"

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type         string        `yaml:"type" env-default:"port"`
		BindIP       string        `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port         string        `yaml:"port" env-default:"8080"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
	Jwt     JwtConfig     `yaml:"jwt"`
	Mail    MailConfig    `yaml:"mail"`
}

type JwtConfig struct {
	AccessTtl  time.Duration `yaml:"access_ttl"`
	RefreshTtl time.Duration `yaml:"refresh_ttl"`
	Secret     string        `yaml:"secret"`
}

type MailConfig struct {
	ServerHostName string `yaml:"server_host_name"`
	ServerPort     string `yaml:"server_port"`
	Account        string `yaml:"account"`
	Password       string `yaml:"password"`
}

type StorageConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	Database        string        `yaml:"database"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	SslMode         string        `yaml:"ssl_mode"`
	MaxOpenCons     int           `yaml:"max_open_cons"`
	MaxIdleCons     int           `yaml:"max_idle_cons"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
