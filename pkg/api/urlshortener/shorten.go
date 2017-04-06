package urlshortener

import (
	"fmt"
	"net/http"
)

// Shorten create a short URL from a long URL
func Shorten(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "test")
}
