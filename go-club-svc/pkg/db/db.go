package db

import (
	"club/pkg/logging"
	"club/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		logger.Error(err)
	}

	db.AutoMigrate(&models.Club{})

	return Handler{
		db,
	}
}