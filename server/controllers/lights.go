package controllers

import (
	"sdk/philips"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

// LightsToggle toggles the lights
func LightsToggle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Room: %v\n", vars["room"])

	err := philips.ToggleRoom(vars["room"])
	if err != nil {
		fmt.Fprint(w, "Error: %v\n", err)
	}
	
}