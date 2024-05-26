package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const setUserActiveQuery = `UPDATE users SET active = true WHERE email = $1 RETURNING active`

type AuthManager struct {
	connection      *sql.DB   // TODO: move all queries to prepared statements and remove connection from struct
	setUserPrepared *sql.Stmt // TODO: suggest possibly sqlx here
}

func NewAuthManager(connection *sql.DB) (*AuthManager, error) {
	setUserPrepared, setUserPreparedErr := connection.Prepare(setUserActiveQuery)
	if setUserPreparedErr != nil {
		return nil, setUserPreparedErr
	}

	return &AuthManager{
		connection,
		setUserPrepared,
	}, nil
}

type LoginCredentials struct {
	EmailID  string `json:"emailID" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserV1 struct {
	ID       uint   `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Active   bool   `json:"active" db:"active"`
}

type UserV2 struct {
	*UserV1
	// Allow null:
	// ActivationDate *time.Time `json:"active" db:"activation_date"`
	// Do not allow null:
	ActivationDate time.Time `json:"active" db:"activation_date"`
}

// To make db changes that are backwards compatible, do one of the following:
// 1. Allow null in the field  - `ALTER TABLE foo ADD activation_date timestamp with time zone not null;`
// 2. Do not allow null but add a default value `ALTER TABLE foo ADD activation_date timestamp with time zone allow null DEFAULT now();`

// Option 1:
func NewUserV1(id uint, email string, username string, password string, active bool) *UserV1 {
	return &UserV1{
		id,
		email,
		username,
		password,
		active,
	}
}

func NewUserV2(id uint, email string, username string, password string, active bool, activationDate time.Time) *UserV2 {
	return &UserV2{
		&UserV1{
			id,
			email,
			username,
			password,
			active,
		},
		activationDate,
	}
}

func (authManager *AuthManager) InsertOption1UserV1(ctx context.Context, user *UserV1) error {
	_, err := authManager.connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active) VALUES($1, $2, $3, $4);
	`, user.Email, user.Username, user.Password, user.Active)
	return err
}

func (authManager *AuthManager) InsertOption1UserV2(ctx context.Context, user *UserV2) error {
	_, err := authManager.connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active, activation_date) VALUES($1, $2, $3, $4);
	`, user.Email, user.Username, user.Password, user.Active, user.ActivationDate)
	return err
}

// Option 2:
func (authManager *AuthManager) InsertOption2UserV1(ctx context.Context, id uint, email string, username string, password string, active bool) error {
	user := &UserV1{
		id,
		email,
		username,
		password,
		active,
	}
	_, err := authManager.connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active ) VALUES($1, $2, $3, $4);
	`, user.Email, user.Username, user.Password, user.Active)
	return err
}

func (authManager *AuthManager) InsertOption2UserV2(ctx context.Context, id uint, email string, username string, password string, active bool, activationDate time.Time) error {
	user := &UserV2{
		&UserV1{
			id,
			email,
			username,
			password,
			active,
		},
		activationDate,
	}
	_, err := authManager.connection.ExecContext(ctx, `
		INSERT INTO users(email, username, password, active, activationDate) VALUES($1, $2, $3, $4, $5);
	`, user.Email, user.Username, user.Password, user.Active, user.ActivationDate)
	return err
}

func (authManager *AuthManager) IsVerified(ctx context.Context, e string) (bool, error) {
	active := new(bool)
	err := authManager.connection.QueryRowContext(ctx, `
		SELECT active FROM users WHERE email = $1;
	`, e).Scan(&active)
	if err != nil {
		return false, err
	}
	return *active, nil
}

func (authManager *AuthManager) Verify(ctx context.Context, m string) error {
	row, rowErr := authManager.setUserPrepared.ExecContext(ctx, m)
	if rowErr != nil {
		return rowErr
	}
	rowsAffected, rowsAffectedErr := row.RowsAffected()
	if rowsAffectedErr != nil {
		return rowsAffectedErr
	}
	if rowsAffected > 0 {
		return fmt.Errorf("id %s does not exist", m)
	}
	return nil
}

func (authManager *AuthManager) fooVerify(ctx context.Context, m string) error {
	if _, err := authManager.setUserPrepared.ExecContext(ctx, m); err != nil {
		return err
	}
	return nil
}

func (authManager *AuthManager) Login(ctx context.Context) (*LoginCredentials, error) {
	loginCredentials := &LoginCredentials{}

	// TODO - you may want to use sqlx, something like this:
	/*
		if authManager.context.QueryRowContextx(ctx, `SELECT email,password FROM users WHERE email = $1;`, loginCredentials.EmailID).StructScan(loginCredentials); err != nil {
			return nil, err
		}
	*/

	if err := authManager.connection.QueryRowContext(ctx, `SELECT email,password FROM users WHERE email = $1;`, loginCredentials.EmailID).Scan(&loginCredentials.EmailID, &loginCredentials.Password); err != nil {
		return nil, err
	}

	return loginCredentials, nil
}
