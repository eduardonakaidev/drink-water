package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
    Name      string    `gorm:"size:100;not null"`
    Email     string    `gorm:"uniqueIndex;not null"`
    Password  string    `gorm:"not null"`
    Profile   Profile   `gorm:"foreignKey:UserID"`
    WaterIntakes  []WaterIntake  `gorm:"foreignKey:UserID"`
}