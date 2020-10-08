package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/data"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userName := data.UserUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
		log.Printf("Json decode error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user = data.User{}
	user.ID = r.Context().Value("user-id").(int)
	user.UserName = userName.Name
	if err := user.UpdateUser(); err != nil {
		log.Printf("Update user name DB error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
