package main

import (
	"log"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/cmd/api"
	"github.com/internal/env"
)

func main() {

	if err := env.InitEnv(); err != nil {
		log.Print("failed to load env file continuing with hard coded defalts")
	}

	log.Print("User server started")

	registerWithEureka()
	api.Init()

}

func registerWithEureka() {

	client := eureka.NewClient([]string{env.GetString("EUREKA_ADDR", "http://localhost:8081/eureka")})

	instance := eureka.NewInstanceInfo(env.GetString("DISCOVERY_ADDR", "localhost:8082"), env.GetString("SERVER_NAME", "user-server"), env.GetString("IP", "127.0.0.1"), env.GetInt("PORT", 8082), uint(env.GetInt("ttl", 30)), false)

	client.RegisterInstance(env.GetString("SERVER_NAME", "user-server"), instance)

	go func() {

		for {

			err := client.SendHeartbeat(instance.App, instance.HostName)

			if err != nil {
				log.Print("Error: Eureka heartbeat failed " + err.Error())
			} else {
				log.Print("Info: Eureka heartbeat success")
			}

			time.Sleep(time.Second * 10)

		}

	}()

}
