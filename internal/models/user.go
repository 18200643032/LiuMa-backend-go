package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id          int64     `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"type:varchar(50) not null"`
	Password    string    `json:"password" gorm:"type:varchar(50) not null"`
	Email       string    `json:"email" gorm:"type:varchar(50) not null"`
	Mobile      string    `json:"mobile" gorm:"unique;not null" binding:"required"`
	LastProject string    `json:"last_project" gorm:"type:varchar(50)"`
	Token       string    `json:"token" gorm:"type:varchar(50) not null"`
	Account     string    `json:"account" gorm:"unique;not null" binding:"required"`
	Status      string    `json:"status" gorm:"unique;not null" binding:"required"`
	CreateTime  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateTime  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "user"
}
