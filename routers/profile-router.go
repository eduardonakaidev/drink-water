package routers

import (
	"github.com/eduardonakaidev/drink-water-api/controllers"
	"github.com/gorilla/mux"
)

func SetupProfileRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/profile/{id}",controllers.GetProfile).Methods("GET")
	router.HandleFunc("/profile",controllers.RegisterProfile).Methods("POST")
	router.HandleFunc("/profile/{id}",controllers.UpdateProfile).Methods("PUT")

	return router
}