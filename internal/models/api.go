package models

import (
	"time"
)

type Api struct {
	ID          string
	Num         int64
	Name        string
	Level       string
	ModuleID    string
	ProjectID   string
	Method      string
	Path        string
	Protocol    string
	DomainSign  string
	Description string
	Header      string
	Body        string
	Query       string
	Rest        string
	CreateTime  time.Time
	UpdateTime  time.Time
	CreateUser  string
	UpdateUser  string
	Status      string
}
