package db

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (database *DatabaseRepo) GetUser(ctx context.Context, filled,filledName string) (*User, error) {

	query := fmt.Sprintf("SELECT * FROM USERS WHERE %s = $1",filledName)

	// query := `SELECT * FROM USERS WHERE id = $1`

	row, err := database.db.QueryContext(ctx, query, filled)

	if err != nil {
		return nil, err
	}

	row.Next()

	var user User

	err = row.Scan(&user.Id, &user.Name, &user.Username, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user,nil

}



