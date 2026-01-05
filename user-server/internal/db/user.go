package db

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id        string
	name      string
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func (database *DatabaseRepo) Insert(ctx context.Context, name, username, password string) error {

	query := `INSERT INTO USERS(name,username,password) VALUES($1,$2,$3)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = database.db.ExecContext(ctx, query, name, username, hashedPassword)

	return err

}

func (database *DatabaseRepo) CheckuserExist(ctx context.Context, username string) bool {

	query := `SELECT EXIST(SELECT 1 USERS WHERE username = $1)`

	var result bool

	database.db.QueryRowContext(ctx, query, username).Scan(&result)

	return result
}
