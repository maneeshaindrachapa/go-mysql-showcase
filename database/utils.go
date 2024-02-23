package database

import (
	"github.com/jinzhu/gorm"
	"github.com/maneeshaindrachapa/go-mysql/models"
)

func InitializeDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
