package adapters

import "net/http"

// Adapter adapt http handler
type Adapter func(http.Handler) http.Handler

// Adapt create a http adapter
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
