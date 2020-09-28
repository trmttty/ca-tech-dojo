package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/api/controller"
)

// Run say hello
func Run() {

	controller.Init()
	http.HandleFunc("/", hello)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
