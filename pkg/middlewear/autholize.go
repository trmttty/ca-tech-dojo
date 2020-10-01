package middlewear

import (
	"net/http"
)

func Authorize(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
