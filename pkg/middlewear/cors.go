package middlewear

import (
	"net/http"
)

func Cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-token")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		h(w, r)
	}
}
