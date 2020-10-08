package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func CreateToken(userID int) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "game-api"
	claims["iat"] = time.Now().Unix()
	claims["sub"] = strconv.Itoa(userID)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ParseToken(tokenString string) (id int, err error) {
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	if !tkn.Valid {
		return
	}
	id, err = strconv.Atoi(claims.Subject)
	return
}

func Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-token")
		if tokenString == "" {
			log.Println("Authentication error")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := ParseToken(tokenString)
		if err != nil {
			log.Println("Authentication error")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user-id", id)
		r = r.WithContext(ctx)
		h(w, r)
	}
}
