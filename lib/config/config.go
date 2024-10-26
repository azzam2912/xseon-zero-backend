package configlib

import (
	"os"
	"xseon-zero/domain/config"
)

func LoadConfig() (*config.Config, error) {
	return loadLocalConfig(), nil
}

// func loadLocalConfig() *config.Config {

// 	return &config.Config{
// 		DB: config.DBConfig{
// 			Host:     "HOST ANJ",
// 			User:     os.Getenv("DB_USER"),
// 			Password: os.Getenv("DB_PASSWORD"),
// 			Name:     os.Getenv("DB_NAME"),
// 			Port:     os.Getenv("DB_PORT"),
// 		},
// 		Service: config.ServiceConfig{
// 			Port: os.Getenv("SERVICE_PORT"),
// 		},
// 	}
// }

func loadLocalConfig() *config.Config {
	return &config.Config{
		DB: config.DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		Service: config.ServiceConfig{
			Port: os.Getenv("SERVICE_PORT"),
		},
	}
}
