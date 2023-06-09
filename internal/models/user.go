package models

import "time"

type User struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Mobile      string    `json:"mobile" gorm:"unique;not null" binding:"required"`
	LastProject string    `json:"last_project"`
	Token       string    `json:"token"`
	Account     string    `json:"account" gorm:"unique;not null" binding:"required"`
	Status      string    `json:"status"`
	CreateTime  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateTime  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "user"
}
