package routers

import (
	"github.com/eduardonakaidev/drink-water-api/controllers"
	"github.com/gorilla/mux"
)

func SetupWaterRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/water-intake/{id}",controllers.GetWatersIntakes).Methods("GET")
	router.HandleFunc("/water-intake",controllers.RegisterWaterIntake).Methods("POST")
	router.HandleFunc("/water-intake/{id}",controllers.UpdateWaterIntake).Methods("PUT")
	router.HandleFunc("/water-intake/{id}",controllers.DeleteWaterInTake).Methods("DELETE")
	return router
}