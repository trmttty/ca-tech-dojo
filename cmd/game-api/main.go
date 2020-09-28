package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/controllers"
)

func init() {
	// get env
}

// Run say hello
func main() {

	http.HandleFunc("/", hello)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, controllers.Test)
}
