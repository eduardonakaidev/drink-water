package repositories

import (
	"github.com/eduardonakaidev/drink-water-api/models"
	"gorm.io/gorm"
)

type WaterIntakeRepository struct {
	DB *gorm.DB
}

// Register cria um novo registro de consumo de água no banco de dados.
func (r *WaterIntakeRepository) Register(waterIntake models.WaterIntake) error {
	if err := r.DB.Create(&waterIntake).Error; err != nil {
		return err
	}
	return nil
}

// GetHistory retorna o histórico de consumo de água de um usuário específico.
func (r *WaterIntakeRepository) GetHistory(userId uint) ([]models.WaterIntake, error) {
	var history []models.WaterIntake

	// Filtra registros pelo ID do usuário
	if err := r.DB.Where("user_id = ?", userId).Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}