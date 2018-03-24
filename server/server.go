package server

import (
	"github.com/gorilla/mux"
	"github.com/ascension/server/controllers"
	"net/http"
)

// Start the server
func Start() error {
	r := mux.NewRouter()

	// Lights routes
	r.HandleFunc("/{room}/lights/toggle", controllers.LightsToggle)
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		return err
	}
	return nil
}