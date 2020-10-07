package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateToken(t *testing.T) {
	id := 1
	_, err := CreateToken(id)
	if err != nil {
		t.Error(err)
	}
}

func TestParseToken(t *testing.T) {
	id := 1
	tokenString, err := CreateToken(id)
	if err != nil {
		t.Error(err)
	}
	ret, err := ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}
	if ret != id {
		t.Error("Wrong id, was expecting 1 but got", ret)
	}
}

func TestAuthenticate(t *testing.T) {
	id := 1
	tokenString, err := CreateToken(id)
	if err != nil {
		t.Error(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", Authenticate(func(w http.ResponseWriter, r *http.Request) {
	}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("x-token", tokenString)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
