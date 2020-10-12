package model

import "testing"

func TestNewUser(t *testing.T) {
	user, err := NewUser("test")
	if err != nil {
		t.Error("failed Test")
	}
	if user.UserName != "test" {
		t.Error("failed Test")
	}
}

func TestSet(t *testing.T) {
	user := User{
		UserName: "test",
	}
	err := user.Set("change")
	if err != nil {
		t.Error("failed Test")
	}
	if user.UserName != "change" {
		t.Error("failed Test")
	}
}
