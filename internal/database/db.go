package database

import (
	"gorm.io/gorm"
	"fmt"
	gmysql "gorm.io/driver/mysql"
	"LiuMa-backend-go/internal/models"
)

var DB *gorm.DB

func InitMysql() { // 1、连接数据库
	dsn := "root:12345678@tcp(localhost:3306)/liuma?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化mysql连接错误", err)
	}
	//DB.SetMaxOpenConns(0)
	DB = db

	if err := DB.AutoMigrate(models.User{}); err != nil {
		fmt.Println("自动创建表结构失败：", err)
		fmt.Println(err)
	}

}
