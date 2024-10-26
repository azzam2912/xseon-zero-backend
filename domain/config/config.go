package config

type Config struct {
	DB               DBConfig               `json:"db"`
	Service          ServiceConfig          `json:"service"`
}

type ServiceConfig struct {
	Port string `json:"port"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     string `json:"port"`
	User     string `json:"user"`
}