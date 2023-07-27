package services

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("jajajjaa")

func GenerateJWT(username string) (string, error) {
	//define new jwt
	token := jwt.New(jwt.SigningMethodHS256)

	//claim jwt
	claim := token.Claims.(jwt.MapClaims)

	//set claim jwt
	claim["username"] = username
	claim["exp"] = time.Now().Add(2 * time.Minute).Unix()

	//signed string jwt
	tokenstr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}

func ParseJWT(tknstring string) (*jwt.Token, error) {
	return jwt.Parse(tknstring, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
