package main

import (
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/internal/auth"
	"github.com/trmttty/ca-tech-dojo/internal/handler"
	mw "github.com/trmttty/ca-tech-dojo/pkg/middlewear"
)

func main() {
	http.HandleFunc("/user/create", mw.Logger(mw.Cors(handler.CreateUser)))
	http.HandleFunc("/user/get", mw.Logger(mw.Cors(auth.Authenticate(handler.GetUser))))
	http.HandleFunc("/user/update", mw.Logger(mw.Cors(auth.Authenticate(handler.UpdateUser))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
