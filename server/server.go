package server

import (
	"github.com/gorilla/mux"
	"github.com/ascension/config"
	"net/http"
)

// Cfg holds all vendor configuration structs.
var Cfg config.VendorConfig

// Start the server
func Start(vCfg config.VendorConfig) error {
	r := mux.NewRouter()
	Cfg = vCfg

	// Lights routes
	r.HandleFunc("/{room}/lights/toggle", LightsToggle)
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		return err
	}
	return nil
}