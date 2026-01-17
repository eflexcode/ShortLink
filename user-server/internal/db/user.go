package db

import (
	"context"
	"fmt"
	"log"
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

func (database *DatabaseRepo) Insert(ctx context.Context, name, username, password string) error {

	query := `INSERT INTO USERS(name,username,password,updated_at) VALUES($1,$2,$3,$4)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = database.db.ExecContext(ctx, query, name, username, hashedPassword,time.Now())

	return err

}

func (database *DatabaseRepo) CheckuserExist(ctx context.Context, username string) bool {

	query := `SELECT EXIST(SELECT 1 USERS WHERE username = $1)`

	var result bool

	database.db.QueryRowContext(ctx, query, username).Scan(&result)

	return result
}

func (database *DatabaseRepo) GetUser(ctx context.Context, filled, filledName string) (*User, error) {

	log.Print(filled+" "+filledName)
	
	query := fmt.Sprintf("SELECT * FROM USERS WHERE %s = $1", filledName)

	c := context.Background()
	
	row, err := database.db.QueryContext(c,query, filled)

	if err != nil {
		log.Print("query failed")
		return nil, err
	}

	defer row.Close()
	
	row.Next()

	var user User

	err = row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Print("scan failed")
		return nil, err
	}

	return &user, nil
}

func (database *DatabaseRepo) Update(ctx context.Context, displayName, id string) error {

	query := `UPDATE USERS SET name = $1 WHERE id = $2`

	_, err := database.db.ExecContext(ctx, query, displayName, id)

	return err
}

func (database *DatabaseRepo) UpdatePassword(ctx context.Context, password, id string) error {

	query := `UPDATE USERS SET password = $1 WHERE id = $2`

	_, err := database.db.ExecContext(ctx, query, password, id)

	return err
}
