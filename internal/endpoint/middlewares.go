package endpoint

import (
	"log"
	"net/http"
	"time"
)

// Logger ...
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
