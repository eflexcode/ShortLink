package db

import (
	"context"
	"database/sql"
	"time"

)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Database struct{
	postgres *sql.DB
}

func(database Database) GetUser(username string,context context.Context) *User{
	
	query := `SELECET * FROM USER WHERE USERNAME = $1`
	
	row,err :=	database.postgres.QueryContext(context,username,query)

	if err != nil{
		return nil
	}	
	
	row.Next()
	var user User
	
	row.Scan(&user.Id,&user.Name,&user.Username,&user.Password,&user.CreatedAt,&user.UpdatedAt)
	
	return &user
	
}