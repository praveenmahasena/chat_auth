package postgres

import (
	"context"

)

type User struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Active   bool   `json:"active" db:"active"`
}

func (u *User) InsertOne(ctx context.Context) error {
	var _, err = connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active ) VALUES($1, $2, $3, $4);
	`, u.Email, u.Username, u.Password, u.Active)
	return err
}

func IsVerified(ctx context.Context, e string) (bool, error) {
	var active = new(bool)
	var err = connection.QueryRowContext(ctx, `
		SELECT active FROM users WHERE email = $1;
	`, e).Scan(&active)
	if err != nil {
		return false, err
	}
	return *active, nil
}
