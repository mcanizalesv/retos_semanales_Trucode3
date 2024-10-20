package models

import (
	"gorm.io/gorm"
)

type Engineer struct {
	gorm.Model
	Name string `json:"name"`
	Surname string `json:"surname"`
	ProjectID  uint `json:"project_id"`
	Project Project
}