package adapters

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/context"
)

// WithDatabase create database adapter
func WithDatabase(db *mgo.Session) Adapter {
	// return the Adapter
	return func(h http.Handler) http.Handler {
		// the adapter (when called) should return a new handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// copy the database session
			dbsession := db.Copy()
			defer dbsession.Close() // clean up
			// save it in the mux context
			context.Set(r, "database", dbsession)
			// pass execution to the original handler
			h.ServeHTTP(w, r)
		})
	}
}
