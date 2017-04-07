package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Index execute the main route in our API
func Index(w http.ResponseWriter, r *http.Request) {

	options := map[string][]string{
		"links": {"/api/v1/shorten", "/api/v1/urls"},
	}
	response, _ := json.Marshal(options)
	fmt.Fprint(w, string(response))
}
