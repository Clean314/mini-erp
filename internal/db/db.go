package db

import (
	"log"
	"mini-erp/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=db user=admin password=admin dbname=mini-erp port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Project{})
}
