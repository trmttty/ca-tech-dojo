package handler

import (
	"log"
	"net/http"

	mw "github.com/trmttty/ca-tech-dojo/pkg/interface/middlewear"
)

func Run(userHandler UserHandler) {
	http.HandleFunc("/user/create", mw.Logger(mw.Cors(userHandler.Post())))
	http.HandleFunc("/user/get", mw.Logger(mw.Cors(mw.Authenticate(userHandler.Get()))))
	http.HandleFunc("/user/update", mw.Logger(mw.Cors(mw.Authenticate(userHandler.Put()))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
