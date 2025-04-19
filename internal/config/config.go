package config

import (
	"log"

	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Config() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func DbConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/main.db"))

	if err != nil {
		log.Fatal("db connection error!", err)
	}

	if err := db.AutoMigrate(&models.Users{}, &models.Books{}, &models.Categories{}); err != nil {
		log.Fatal("db table creation error!", err)
	}

	return db
}
