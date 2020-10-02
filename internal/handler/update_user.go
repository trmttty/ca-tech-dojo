package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/internal/auth"
	"github.com/trmttty/ca-tech-dojo/internal/data"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("x-token")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = data.User{}
	var err error
	user.ID, err = auth.ParseToken(tokenString)
	if err != nil {
		log.Printf("Parse token error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.UpdateUser()
	if err != nil {
		log.Printf("Update user name DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
