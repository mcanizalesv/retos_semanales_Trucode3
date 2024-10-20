package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name          string
	Delivery_Date time.Time
	Engineers     []Engineer  // Project has many Engineers
    Tasks         []Task  // Project has many Tasks
}