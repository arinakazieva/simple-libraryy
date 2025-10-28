package config

import "errors"

type Config struct {
	DatabaseURL string
	Port        string
	Debug       bool
}

func New() *Config {
	return &Config{
		DatabaseURL: "localhost:5432",
		Port:        "8080",
		Debug:       true,
	}
}

// Validate проверяет корректность конфигурации
func (c *Config) Validate() error {
	if c.DatabaseURL == "" {
		return errors.New("database URL cannot be empty") // исправлено: с маленькой буквы
	}
	if c.Port == "" {
		return errors.New("port cannot be empty") // исправлено: с маленькой буквы
	}
	return nil
}