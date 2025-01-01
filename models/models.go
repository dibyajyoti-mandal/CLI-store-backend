package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name      string
	Price     int
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
