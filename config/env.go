package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environments struct {
	JWT_SECRET string
	DB_URI     string
	PORT       int64
}
// Configura o env
func Env() (*Environments) {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtém as variáveis
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URI := os.Getenv("DB_URI")
	PORT_STR := os.Getenv("PORT")
	// Validações
	if JWT_SECRET == "" {
		log.Fatalf("JWT_SECRET is required and not set in the .env file")
	}
	if DB_URI == "" {
		log.Fatalf("DB_URI is required and not set in the .env file")
	}
	if PORT_STR == "" {
		log.Fatalf("PORT is required and not set in the .env file")
	}

	// Converte o PORT para inteiro
	PORT, err := strconv.ParseInt(PORT_STR, 10, 64)
	if err != nil {
		log.Fatalf("PORT must be a valid integer")
	}

	// Retorna as variáveis de ambiente encapsuladas na struct
	return &Environments{
		JWT_SECRET: JWT_SECRET,
		DB_URI:     DB_URI,
		PORT:       PORT,
	}
}
