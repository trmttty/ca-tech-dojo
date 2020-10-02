package middlewear

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h(w, r)
		endTime := time.Since(startTime)
		log.Printf("%s %s %s %v", r.URL, r.Method, r.Host, endTime)
	}
}
