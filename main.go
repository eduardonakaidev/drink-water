package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardonakaidev/drink-water-api/config"
	database "github.com/eduardonakaidev/drink-water-api/config/db"
	"github.com/eduardonakaidev/drink-water-api/routers"
	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o banco de dados
	database.InitDatabase()

	// Inicia as migrations
	database.Migrate()

	// Configura as rotas
	router := mux.NewRouter()
	routers.SetupAuthicationRoutes(router)
	routers.SetupProfileRoutes(router) 
	routers.SetupWaterRoutes(router)   
	routers.SetupHearthRoutes(router)   

	
	env_ := config.Env()
	// Define a porta do servidor
	address := fmt.Sprintf(":%d", env_.PORT)
	// Inicia o servidor...
	log.Printf("Servidor rodando na porta %v...",address)
	log.Fatal(http.ListenAndServe(address, router))
}
