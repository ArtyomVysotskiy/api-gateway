package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type (
	// Config - структура конфига проекта
	Config struct {
		App           AppConfig           `yaml:"app"`
		Microservices MicroservicesConfig `yaml:"microservices"`
	}
	// AppConfig - структура конфига приложения
	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Port    int    `yaml:"port"`
	}
	// MicroservicesConfig - структура конфига микросервисов
	MicroservicesConfig struct {
		AuthSvcUrl string `yaml:"authSvcUrl"`
	}
)

// NewConfig - конструктор для создания Config
func NewConfig() (*Config, error) {
	// Создаем конфигурацию
	cfg := &Config{}
	// Загружаем конфигурацию с использованием cleanenv
	if err := cleanenv.ReadConfig("./config/config.yaml", cfg); err != nil {
		log.Println("Error loading environment variables:", err)
		return nil, err
	}
	return cfg, nil
}
