package adapters

import (
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/souzavitor/go-study/pkg/api/helpers"
)

// WithDatabase create database adapter
func WithDatabase() Adapter {
	session := helpers.CreateDatabaseConnection()
	// return the Adapter
	return func(h http.Handler) http.Handler {
		// the adapter (when called) should return a new handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// copy session
			reqSession := session.Clone()
			defer reqSession.Close()
			db := reqSession.DB(os.Getenv("DB_DATABASE"))

			// save it in the mux context
			context.Set(r, "database", db)

			// pass execution to the original handler
			h.ServeHTTP(w, r)
		})
	}
}
