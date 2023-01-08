package config

import "time"

const (
	UnixSockType   = "sock"
	TcpListener    = "tcp"
	UnixListener   = "unix"
	DefaultTimeout = 1 * time.Second
)
