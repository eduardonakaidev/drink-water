package routers

import (
	"github.com/eduardonakaidev/drink-water-api/controllers"
	"github.com/gorilla/mux"
)

func SetupAuthicationRoutes(router *mux.Router) *mux.Router {
	
   
	router.HandleFunc("/register",controllers.RegisterUser).Methods("POST")

	router.HandleFunc("/login",controllers.LoginUser).Methods("POST")
	
	router.HandleFunc("/refresh",controllers.RefreshUser).Methods("POST")
	
	router.HandleFunc("/logout",controllers.LogoutUser).Methods("POST")

	return router
}