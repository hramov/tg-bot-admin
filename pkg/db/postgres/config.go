package postgres

import "time"

type Config struct {
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
