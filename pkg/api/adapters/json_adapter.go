package adapters

import (
	"net/http"
)

// WithContentTypeJson set header content type as application/json
func WithContentTypeJson() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	}
}
