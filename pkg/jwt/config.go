package jwt

import "time"

type Config struct {
	AccessTtl     time.Duration `yaml:"access_ttl"`
	RefreshTtl    time.Duration `yaml:"refresh_ttl"`
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
}
