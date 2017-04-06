package api

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/souzavitor/go-study/pkg/api/adapters"
)

// NewRouter create a new object to handle routes
func NewRouter() *mux.Router {
	db, err := mgo.Dial(os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = adapters.Adapt(
			handler,
			adapters.WithDatabase(db),
			adapters.WithLogger(handler, route.Name),
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
