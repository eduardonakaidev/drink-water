package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID        uint      `gorm:"not null;uniqueIndex"` // Relacionado ao ID do usuário
    Weight        float64   `gorm:"not null"`            // Peso do usuário em kg
    ActivityLevel string    `gorm:"size:20;not null"`    // Nível de atividade física ("low", "moderate", "high")
    DailyGoal     float64   `gorm:"not null"`            // Meta diária de consumo (em mL)
}