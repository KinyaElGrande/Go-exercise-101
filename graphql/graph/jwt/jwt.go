package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("graphqlSecreT")

//GenerateToken generates a jwt token and assigns username to its claims then returns it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}

	return tokenString, nil
}

//ParseToken is going to be used whenever we receive a token and want to know who sent this token
func ParseToken(tokenStr string) (string, error) {
	token , err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
