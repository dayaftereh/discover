package http

import (
	"time"
)

type Config struct {
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Port:         4000,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
