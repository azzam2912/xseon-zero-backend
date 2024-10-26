package db

import (
	"fmt"
	"xseon-zero/domain/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.Name, config.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
