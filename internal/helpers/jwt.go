package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key    []byte = []byte("yayayayaa")
	verify []byte = []byte("verify")
)

func GenerateJWT(m string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Email": m,                                          // TODO: change the "Email" to const
			"exp":   time.Now().Add(time.Hour * 24 * 15).Unix(), // TODO: change the "exp" to const
		},
	)
	return token.SignedString(key)
}

func DecodeJWT(t string) (string, error) {
	token, tokenErr := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if tokenErr != nil {
		return "", tokenErr
	}
	if !token.Valid {
		return "", errors.New("Invalid Token")
	}
	return token.Claims.(jwt.MapClaims)["Email"].(string), nil // TODO: change the "Email" to const
}

func GenerateJWTVerify(m string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": m, // // TODO: change the "email" to const
	})
	return token.SignedString(verify)
}

func DecodeJWTVerify(m string) (string, error) {
	token, tokenErr := jwt.Parse(m, func(t *jwt.Token) (interface{}, error) {
		return verify, nil
	})
	if tokenErr != nil {
		return "", tokenErr
	}
	if !token.Valid {
		return "", fmt.Errorf("token not Valid")
	}
	return token.Claims.(jwt.MapClaims)["email"].(string), nil // TODO: change the "email" to const
}
