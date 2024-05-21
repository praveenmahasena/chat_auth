package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(b string) (string, error) {
	p, e := bcrypt.GenerateFromPassword([]byte(b), 10)
	return string(p), e
}
