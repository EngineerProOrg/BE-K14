package config

import "time"

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	ServerPort    string
	RateLimit     struct {
		Requests int
		Window   time.Duration
	}
}

func NewConfig() *Config {
	return &Config{
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
		ServerPort:    ":8080",
		RateLimit: struct {
			Requests int
			Window   time.Duration
		}{
			Requests: 2,                // 2 requests allowed
			Window:   60 * time.Second, // per 60 seconds
		},
	}
}
