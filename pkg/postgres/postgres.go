package postgres

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"url-shortener/internal/config"
	"url-shortener/internal/models"
)

func NewInstance() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString(config.DatabaseHost),
		viper.GetString(config.DatabasePort),
		viper.GetString(config.DatabaseUsername),
		viper.GetString(config.DatabasePassword),
		viper.GetString(config.DatabaseName),
		viper.GetString(config.DatabaseSslMode),
	)))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&models.Link{}); err != nil {
		log.Fatalf("%v", err)
	}

	return db, nil
}
