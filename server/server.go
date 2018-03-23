package server

import (
	"github.com/gorilla/mux"
	"github.com/ascension/server/controllers"
	"net/http"
)

// Start the server
func Start() {
	r := mux.NewRouter()

	// Lights routes
	r.HandleFunc("/{room}/lights/toggle", controllers.LightsToggle)
	http.ListenAndServe(":10000", r)
}