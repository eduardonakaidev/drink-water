package controllers

import (
	"net/http"

	"github.com/eduardonakaidev/drink-water-api/views"
)

func HearthController(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"message": "Hello World",
	}
	views.RenderJSON(w,response , http.StatusOK)
}
