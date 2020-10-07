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

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	userName := data.UserCreateRequest{}
	err := json.Unmarshal(body, &userName)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := data.User{}
	user.UserName = userName.Name
	err = user.CreateUser()
	if err != nil {
		log.Printf("Create user DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := data.UserCreateResponse{}
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
