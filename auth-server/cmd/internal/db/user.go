package db

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegUser struct {
	DisplayName string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Database struct {
	Postgres *sql.DB
}

func (database Database) GetUser(username string, context context.Context) *User {

	query := `SELECT * FROM USERS WHERE USERNAME = $1`

	row, err := database.Postgres.QueryContext(context, query, username)

	if err != nil {
		println(err.Error())
		return nil
	}

	row.Next()
	var user User

	row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user
}

func (database Database) ResetPassword(username, password string, context context.Context) error {

	query := `UPDATE users SET password = $1 WHERE username $2`

	_, err := database.Postgres.ExecContext(context, query, password, username)

	return err
}

func (database Database) Register(regUser RegUser, context context.Context) error {

	createdAt := time.Now()

	query := `INSERT INTO users(name,username,password,created_at,updated_at) VALUES($1,$2,$3,$4,$5)`

	password, err := bcrypt.GenerateFromPassword([]byte(regUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = database.Postgres.ExecContext(context, query, regUser.DisplayName, regUser.Username, password, createdAt, createdAt)

	return err
}
