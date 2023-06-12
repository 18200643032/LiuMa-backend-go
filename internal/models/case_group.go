package models

import (
	"gorm.io/gorm"
)

// CasesGroup 表
type CasesGroup struct {
	gorm.Model
	//ProjectID int64
	Name   string `gorm:"name"`
	Parent int64  `gorm:"parent"`
	Child  int64  `gorm:"child"`
}

func (CasesGroup) TableName() string {
	return "case_group"
}
