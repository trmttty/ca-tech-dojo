package auth

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create jwt token
func CreateToken(usersID int) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "game-api"
	claims["sub"] = strconv.Itoa(usersID)
	// claims["iat"] = time.Now().Format("2006-01-02 15:04:05")
	tokenString, err = token.SignedString([]byte("my_secret_key"))
	return
}

type Credentials struct {
	Username string `json:"name"`
}

//
// type Claims struct {
// 	UsersID string `json:"id"`
// 	jwt.StandardClaims
// }

var jwtKey = []byte("my_secret_key")

func ParseToken(tokenString string) (id int, err error) {
	claims := &jwt.StandardClaims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	if !tkn.Valid {
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	id, err = strconv.Atoi(claims.Subject)
	return
}
