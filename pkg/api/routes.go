package api

import (
	"net/http"

	"github.com/souzavitor/go-study/pkg/api/handlers"
)

// Route struct to specify how a route must be
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes new type of array of Routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		handlers.Index,
	},
	Route{
		"ShortenURL",
		"POST",
		"/api/v1/shorten",
		handlers.Shorten,
	},
}
