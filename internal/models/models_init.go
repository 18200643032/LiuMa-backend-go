package models

import (
	"LiuMa-backend-go/internal/database"
	"log"
)

func Create(mode interface{}) {
	if err := database.DB.Set("gorm:table_options",
		"charset=utf8mb4").AutoMigrate(
		mode); err != nil {
		log.Printf("Create Table :%s\n", err.Error())
	}
}
