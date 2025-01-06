package routers

import (
	"github.com/eduardonakaidev/drink-water-api/controllers"
	"github.com/gorilla/mux"
)

func SetupHearthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/",controllers.HearthController).Methods("GET")

	return router
}
