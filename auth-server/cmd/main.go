package main

import (
	"log"

	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/cmd/env"
	"github.com/cmd/internal/db"
	"github.com/cmd/api"
)

func main(){
	
	log.Print("Auth server started")
	env.InitEnv()
	
 	config := db.ConfigDb{
		DbType:       env.GetString("DB_TYPE", "postgres"),
		Addr:         env.GetString("DB_ADDR", "postgres://postgres:12345@localhost/shortlinkuser?sslmode=disable"),
		MaxOpenConn:  env.GetInt("MAX_OPEN_CONN", 20),
		MaxIdealConn: env.GetInt("MAX_IDEA_CONN", 20),
		MaxIdealTime: env.GetString("MAX_IDEAL_TIME", "15m"),
	}

 	db,err := db.DatabaseConnect(config)
  
  	if	err != nil {
 		panic("Failed to connect to db")
   }
	
	registerWithEureka()
	
	api.InitApi(db)

}

func registerWithEureka() {

	client := eureka.NewClient([]string{"http://localhost:8081/eureka"})

	instance := eureka.NewInstanceInfo("localhost:8084", "auth-server", "127.0.0.1", 8084, 30, false)

	client.RegisterInstance("auth-server", instance)

	go func() {

		for {
			
			err := client.SendHeartbeat(instance.App, instance.HostName)

			if err != nil {
				log.Print("Error: Eureka heartbeat failed ")
			}else {
				log.Print("Info: Eureka heartbeat success")
			}

			time.Sleep(time.Second * 100)

		}

	}()

}


