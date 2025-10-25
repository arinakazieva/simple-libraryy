package main

import (
	"errors"
)

func GetPortFromConfig(cfg map[string]string) (string, error) {
	port, exists := cfg["PORT"]
	if !exists {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}
