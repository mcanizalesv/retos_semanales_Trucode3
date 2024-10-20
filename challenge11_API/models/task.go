package models

import (
	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
    Title           string `json:"title"`
    Description     string `json:"description"` 
    EstimatedHours  int    `json:"estimated_hours"` 
    ProjectID       uint   `json:"project_id"`
}