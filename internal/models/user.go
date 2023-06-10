package models

import (
	"LiuMa-backend-go/internal/database"
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
	// 建表
	Create(&User{})
	return "user"
}

type UserLoginCommand struct {
	Username string `json:"username" gorm:"type:varchar(50) not null"`
	Password string `json:"password" gorm:"type:varchar(50) not null"`
}

type UserRegisterCommand struct {
	Username string `json:"username" gorm:"type:varchar(50) not null"`
	Email    string `json:"email" gorm:"type:varchar(50) not null"`
	Mobile   string `json:"mobile" gorm:"unique;not null" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(50) not null"`
}

// 获取单个用户信息
func (model *User) UserGet(id int) (user User, err error) {
	err = database.DB.Where("id = ?", id).Find(&user).Error
	return
}

// 获取用户列表
func (model *User) UserGetList(page int, pageSize int) (users []User, err error) {
	offset := pageSize * (page - 1)
	err = database.DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return
}

// 创建用户
func (model *User) UserAdd() (id int64, err error) {
	err = database.DB.Create(model).Error
	return
}

// 更新用户
func (model *User) UserUpdate(id int) (err error) {
	err = database.DB.Model(model).Where("id = ?", id).Updates(User{Username: model.Username, Password: model.Password}).Error
	return
}

// 删除用户
func (model *User) UserDelete(id int) (err error) {
	err = database.DB.Where("id = ?", id).Delete(User{}).Error
	return
}
