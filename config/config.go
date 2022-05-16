package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
}

var appConfig *AppConfig

func AddConfig() *AppConfig {

	if appConfig == nil {
		appConfig = InitConfig()
	}

	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "layered_db")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "")

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
