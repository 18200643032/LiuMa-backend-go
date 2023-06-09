package models

import "time"

type Project struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ProjectAdmin string    `json:"project_admin"`
	CreateTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Status       string    `json:"status"`
}

func (Project) TableName() string {
	return "project"
}
