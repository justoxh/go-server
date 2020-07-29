package config

import (
	"sync"

	"github.com/jinzhu/configor"
)

var (
	cfg Configuration
	mu  sync.RWMutex
)

func Load(file *string) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()
	err := configor.Load(&cfg, *file)
	if err != nil {
		return Configuration{}, err
	}
	return cfg, err
}

func GetConfig() Configuration {
	mu.Lock()
	defer mu.Unlock()
	return cfg
}
