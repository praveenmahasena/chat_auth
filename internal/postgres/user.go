package postgres

import "context"

type User struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Active   bool   `json:"active" db:"active"`
}

func NewUser() User {
	return User{}
}

func (u *User) Check(ctx context.Context) bool {

}
