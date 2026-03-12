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

type RegUser struct{
	DisplayName string `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

type Database struct{
	Postgres *sql.DB
}

func(database Database) GetUser(username string,context context.Context) *User{
	
	query := `SELECT * FROM USER WHERE USERNAME = $1`
	
	row,err :=	database.Postgres.QueryContext(context,username,query)

	if err != nil{
		return nil
	}	
	
	row.Next()
	var user User
	
	row.Scan(&user.Id,&user.Name,&user.Username,&user.Password,&user.CreatedAt,&user.UpdatedAt)
	
	return &user
}

func(database Database) Register(regUser RegUser,context context.Context) error{
	
	createdAt := time.Now()
	
	query := `INSERT INTO user(name,username,password,created_at,updated_at)`
	
	password,err  := bcrypt.GenerateFromPassword([]byte(regUser.Password),bcrypt.DefaultCost)
	
	if err != nil{
		return err
	}
	
	_,err = database.Postgres.ExecContext(context,query,regUser.DisplayName,regUser.Username,password,createdAt,createdAt)
	
	return err
}