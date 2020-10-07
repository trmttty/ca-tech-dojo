package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/auth"
	"github.com/trmttty/ca-tech-dojo/pkg/data"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("x-token")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := auth.ParseToken(tokenString)
	if err != nil {
		log.Printf("Parse token error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userName := data.UserGetResponse{}
	userName.Name, err = data.GetUserName(id)
	if err != nil {
		log.Printf("Get user name DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.MarshalIndent(&userName, "", " ")
	if err != nil {
		log.Printf("Encode token error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
