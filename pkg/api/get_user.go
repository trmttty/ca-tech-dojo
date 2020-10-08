package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/pkg/data"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value("user-id").(int)

	var err error
	userName := data.UserGetResponse{}
	userName.Name, err = data.GetUserName(userID)
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
