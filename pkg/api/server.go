package api

import (
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/auth"
	mw "github.com/trmttty/ca-tech-dojo/pkg/middlewear"
)

func Run() {
	http.HandleFunc("/user/create", mw.Logger(mw.Cors(CreateUser)))
	http.HandleFunc("/user/get", mw.Logger(mw.Cors(auth.Authenticate(GetUser))))
	http.HandleFunc("/user/update", mw.Logger(mw.Cors(auth.Authenticate(UpdateUser))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
