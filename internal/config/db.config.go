package config

import "time"

const MaxOpenConns = 10
const MaxIdleConns = 5
const ConnMaxIdleTime = 30 * time.Second
const ConnMaxLifetime = 1 * time.Minute