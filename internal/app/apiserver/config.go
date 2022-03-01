package apiserver

import "github.com/FreezeOMatic/Firsttry/internal/app/store"

// .config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log-level"`
	Store    *store.Config
}

// New config ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
