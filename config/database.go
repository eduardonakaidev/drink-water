package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Inicializa a conexão com o banco de dados
func InitDatabase() {
	var err error
	dsn := "host=localhost user=adminuser password=adminpassword dbname=postgresdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
}