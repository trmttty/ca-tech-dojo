package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/internal/auth"
	"github.com/trmttty/ca-tech-dojo/internal/data"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	user := data.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.CreateUser()
	if err != nil {
		log.Printf("Create user DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var token = make(map[string]interface{})
	token["token"], err = auth.CreateToken(user.ID)
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
