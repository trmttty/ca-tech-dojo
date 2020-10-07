package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/auth"
	"github.com/trmttty/ca-tech-dojo/pkg/data"
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

	userName := data.UserUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
		log.Printf("Json decode error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.UserName = userName.Name
	err = user.UpdateUser()
	if err != nil {
		log.Printf("Update user name DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
