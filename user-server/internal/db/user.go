package db

import "time"

type User struct{
	id string
	name string
	username string
	password string
	createdAt time.Time
	updatedAt time.Time
}

func Insert (name,username,password string) error{

}

func ( database *DatabaseRepo) checkuserExist(username string) bool{

database.db.

}