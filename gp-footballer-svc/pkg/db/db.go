package db

import (
	"football/pkg/logging"
	"football/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		logger.Error(err)
	}

	db.AutoMigrate(&models.Footballer{})

	return Handler{
		DB: db,
	}
}