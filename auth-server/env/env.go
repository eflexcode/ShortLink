package env

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

func InitEnv() error{
	return godotenv.Load()
}

func GetInt(key string,fallback  int )int{
	
	env := os.Getenv(key)
	
	int,err := strconv.Atoi(env)
	
	if err != nil{
		return  fallback
	}
	return   int
}

func GetString(key ,fallback  string )string{
	
	env := os.Getenv(key)
	
	if env == ""{
		return  fallback
	}
	
	return  env
}