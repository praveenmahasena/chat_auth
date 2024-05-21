package helpers

import (
	"errors"
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
			"Email": m,
			"exp":   time.Now().Add(time.Hour * 24 * 15).Unix(),
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
	return token.Claims.(jwt.MapClaims)["Email"].(string), nil
}

func GenerateJWTVerify(m string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": m,
	})
	return token.SignedString(verify)
}
