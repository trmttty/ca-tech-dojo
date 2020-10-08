package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/auth"
	"github.com/trmttty/ca-tech-dojo/pkg/data"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userName := data.UserCreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
		log.Printf("Json decode error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := data.User{}
	user.UserName = userName.Name
	if err := user.CreateUser(); err != nil {
		log.Printf("Create user DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := data.UserCreateResponse{}
	var err error
	token.Token, err = auth.CreateToken(user.ID)
	if err != nil {
		log.Printf("Create token error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.MarshalIndent(&token, "", " ")
	if err != nil {
		log.Printf("Encode token error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
	return
}
