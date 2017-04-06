package adapters

import (
	"log"
	"net/http"
	"time"
)

// WithLogger create adapter to handle to log access
func WithLogger(inner http.Handler, name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			inner.ServeHTTP(w, r)

			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}
