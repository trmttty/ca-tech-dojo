package models

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"name"`
	Token     Token
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser create user
func (user *User) CreateUser(db *sql.DB) (err error) {
	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", user.UserName)
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

func GetUserName(id int, db *sql.DB) (userName string, err error) {
	err = db.QueryRow("SELECT name from users WHERE id = ?", id).Scan(&userName)
	return
}

func (user *User) UpdateUser(db *sql.DB) (err error) {
	_, err = db.Exec("UPDATE users SET name=? WHERE id=?", user.UserName, user.ID)
	return
}
