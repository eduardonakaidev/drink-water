package repositories

import (
	"errors"
	"time"

	"github.com/eduardonakaidev/drink-water-api/models"
	"gorm.io/gorm"
)

type WaterIntakeRepository struct {
	DB *gorm.DB
}

// Register cria um novo registro de consumo de água no banco de dados.
func (r *WaterIntakeRepository) Register(waterIntake models.WaterIntake) (models.WaterIntake, error) {
	if err := r.DB.Create(&waterIntake).Error; err != nil {
		return waterIntake, err
	}
	return waterIntake, nil
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

// UpdateWaterIntake atualiza o registro de consumo de água de um usuário.
func (r *WaterIntakeRepository) UpdateWaterIntake(waterIntake models.WaterIntake, userId uint) (models.WaterIntake, error) {
	// Verifica se o registro existe
	var existing models.WaterIntake
	if err := r.DB.Where("id = ? AND user_id = ?", waterIntake.ID, userId).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return waterIntake, errors.New("registro de consumo não encontrado")
		}
		return waterIntake, err
	}

	// Verifica se o registro é de hoje
	today := time.Now().Truncate(24 * time.Hour) // Obtém o início do dia atual
	recordDate := existing.Timestamp.Truncate(24 * time.Hour)

	if !recordDate.Equal(today) {
		return waterIntake, errors.New("só é permitido atualizar registros do dia atual")
	}

	// Atualiza o registro
	if err := r.DB.Model(&existing).Updates(waterIntake).Error; err != nil {
		return waterIntake, err
	}

	return existing, nil
}