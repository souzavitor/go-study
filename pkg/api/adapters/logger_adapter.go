package adapters

import (
	"log"
	"net/http"
	"time"
)

// WithLogger create adapter to handle to log access
func WithLogger(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get the request start time
			start := time.Now()

			// execute the main handler
			h.ServeHTTP(w, r)

			// log result after the request is finished
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
