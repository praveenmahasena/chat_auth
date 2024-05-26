package postgres

import (
	"context"
)

type LoginCredentials struct {
	EmailID  string `json:"emailID" db:"email"`
	Password string `json:"password" db:"password"`
}

type User struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Active   bool   `json:"active" db:"active"`
}

func NewUserStruct() *User {
	return &User{}
}

func (u *User) InsertOne(ctx context.Context) error {
	_, err := connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active ) VALUES($1, $2, $3, $4);
	`, u.Email, u.Username, u.Password, u.Active)
	return err
}

func IsVerified(ctx context.Context, e string) (bool, error) {
	active := new(bool)
	err := connection.QueryRowContext(ctx, `
		SELECT active FROM users WHERE email = $1;
	`, e).Scan(&active)
	if err != nil {
		return false, err
	}
	return *active, nil
}

func Verify(ctx context.Context, m string) (bool, error) {
	var r bool
	err := connection.QueryRowContext(ctx, `
		UPDATE users SET active = true WHERE email = $1 RETURNING active;
	`, m).Scan(&r)
	return r, err
}

func NewLoginCredentials() *LoginCredentials {
	return &LoginCredentials{}
}

func (l *LoginCredentials) Login(ctx context.Context) (LoginCredentials, error) {
	c := NewLoginCredentials() // I hate doing this
	err := connection.QueryRowContext(ctx, `
		SELECT email,password FROM users WHERE email = $1;
	`, l.EmailID).Scan(&c.EmailID, &c.Password)
	return *c, err
}
