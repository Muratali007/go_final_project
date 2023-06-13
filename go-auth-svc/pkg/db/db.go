package db

import (
	"auth/pkg/logging"
	"auth/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var logger = logging.GetLogger()

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db,err := gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		logger.Error(err)
	}

	db.AutoMigrate(&models.User{})
	return Handler{db}
}