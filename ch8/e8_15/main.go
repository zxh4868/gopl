package main

import (
	"sync"
)

type Config struct{}

var instance *Config

var once sync.Once

func InitConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}
