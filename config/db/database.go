package database

import (
	"log"

	"github.com/eduardonakaidev/drink-water-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Inicializa a conexão com o banco de dados
func InitDatabase() {
	env_ := config.Env()
	DB, err = gorm.Open(postgres.Open(env_.DB_URI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
}