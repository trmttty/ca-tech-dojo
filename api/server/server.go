package server

import (
	"fmt"
	"log"
	"net/http"
)

// Run say hello
func Run() {
	http.HandleFunc("/", hello)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
