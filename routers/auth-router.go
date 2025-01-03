package routers

import (
	"github.com/eduardonakaidev/drink-water-api/controllers"
	"github.com/gorilla/mux"
)

func SetupAuthicationRoutes() *mux.Router {
	router := mux.NewRouter()
   
	router.HandleFunc("/register",controllers.Register).Methods("POST")

	router.HandleFunc("/login",controllers.Login).Methods("POST")
	
	router.HandleFunc("/refresh",controllers.Refresh).Methods("POST")
	
	router.HandleFunc("/logout",controllers.Logout).Methods("POST")

	return router
}