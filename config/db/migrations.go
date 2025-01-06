package database

import "github.com/eduardonakaidev/drink-water-api/models"

// Realiza as migrações
func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.WaterIntake{},
	)
}
