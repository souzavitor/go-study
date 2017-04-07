package adapters

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/context"
)

// WithDatabase create database adapter
func WithDatabase() Adapter {
	session, err := mgo.Dial(os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
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

			log.Println("database adapter")
			// pass execution to the original handler
			h.ServeHTTP(w, r)
		})
	}
}
