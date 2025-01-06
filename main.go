package main

import (
	database "github.com/eduardonakaidev/drink-water-api/config"
)

func main() {
	// Inicializa o banco de dados
	database.InitDatabase()

	// Inicia as migrations
	database.Migrate()
}
