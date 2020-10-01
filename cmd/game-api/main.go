package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/internal/models"
	"github.com/trmttty/ca-tech-dojo/pkg/auth"
	mw "github.com/trmttty/ca-tech-dojo/pkg/middlewear"

	_ "github.com/go-sql-driver/mysql"
)

// Db id database handler
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "game:game@tcp(game_db:3306)/game_db?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
}

//////////////////////////

// Run say hello
func main() {

	http.HandleFunc("/user/create", mw.Cors(createUser))
	http.HandleFunc("/user/get", mw.Cors(getUser))
	http.HandleFunc("/user/update", mw.Cors(updateUser))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	user := models.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	err = user.CreateUser(Db)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	token := models.Token{}
	token.Value, err = auth.CreateToken(user.ID)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	response, err := json.MarshalIndent(&token, "", " ")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(response)
	return
}

func getUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("x-token")
	id, err := auth.ParseToken(tokenString)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var userName = make(map[string]interface{})
	userName["name"], err = models.GetUserName(id, Db)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	response, err := json.MarshalIndent(&userName, "", " ")
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	var err error
	tokenString := r.Header.Get("x-token")
	user.ID, err = auth.ParseToken(tokenString)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	err = user.UpdateUser(Db)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
