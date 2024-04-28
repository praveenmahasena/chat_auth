package postgres

import "context"

type User struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func NewUser() User {
	return User{}
}

func (u *User) Insert(ctx context.Context) error {
	var _, err = connection.ExecContext(ctx, `
		INSERT INTO users (email, username, password) VALUES ($1,$2,$3) RETURNING id;
	`, u.Email, u.Username, u.Password)
	return err
}
