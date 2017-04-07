package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// indexResponse interface to create a response of this route
type indexResponse struct {
	Version string   `json:"version"`
	Links   []string `json:"links"`
}

// Index execute the main route in our API
func Index(w http.ResponseWriter, r *http.Request) {
	options := &indexResponse{
		Version: "v1",
		Links:   []string{"/api/v1/shorten", "/api/v1/urls"},
	}

	response, _ := json.Marshal(options)

	fmt.Fprint(w, string(response))
}
