package model

import (
	"errors"
	"time"
)

// User model
type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("enter name")
	}

	user := &User{
		UserName: name,
	}

	return user, nil
}

func (user *User) Set(name string) error {
	if name == "" {
		return errors.New("enter name")
	}

	user.UserName = name

	return nil
}
