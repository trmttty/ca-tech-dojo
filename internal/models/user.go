package models

import (
	"time"

	db "github.com/trmttty/ca-tech-dojo/internal/database"
)

// User model
type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) CreateUser() (err error) {
	res, err := db.Db.Exec("INSERT INTO users (name) VALUES (?)", user.UserName)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	user.ID = int(id)
	return
}

func (user *User) UpdateUser() (err error) {
	_, err = db.Db.Exec("UPDATE users SET name=? WHERE id=?", user.UserName, user.ID)
	return
}

func GetUserName(id int) (userName string, err error) {
	err = db.Db.QueryRow("SELECT name from users WHERE id = ?", id).Scan(&userName)
	return
}
