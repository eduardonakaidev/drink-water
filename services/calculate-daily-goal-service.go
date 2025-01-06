package services

import (
	"fmt"
	"github.com/eduardonakaidev/drink-water-api/models"
)

func CalculateDailyGoalService(profile *models.Profile) error {
	// Validação do peso
	if profile.Weight <= 0 {
		return fmt.Errorf("weight must be greater than zero")
	}

	// Fórmula base
	baseGoal := 35.0 * profile.Weight

	// Ajustes baseados no nível de atividade
	switch profile.ActivityLevel {
	case "low":
		profile.DailyGoal = baseGoal
	case "moderate":
		profile.DailyGoal = baseGoal * 1.1 // +10%
	case "high":
		profile.DailyGoal = baseGoal * 1.2 // +20%
	default:
		return fmt.Errorf("invalid activity level: must be 'low', 'moderate', or 'high'")
	}

	return nil
}
