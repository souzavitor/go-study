package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souzavitor/go-study/pkg/api/adapters"
)

// NewRouter create a new object to handle routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = adapters.Adapt(
			route.HandlerFunc,
			adapters.WithDatabase(),
			adapters.WithLogger(route.Name),
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
