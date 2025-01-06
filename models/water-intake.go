package models

import (
	"time"

	"gorm.io/gorm"
)

type WaterIntake struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
    Amount    float64   `gorm:"not null"` // Quantidade em mL
    Timestamp time.Time `gorm:"not null"`
}