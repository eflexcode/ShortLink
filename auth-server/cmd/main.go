package main

import (
	"log"

	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"

)

func main(){
	
	log.Print("Auth server started")

	registerWithEureka()


}

func registerWithEureka() {

	client := eureka.NewClient([]string{"http://localhost:8081/eureka"})

	instance := eureka.NewInstanceInfo("localhost:8083", "auth-server", "127.0.0.1", 8083, 30, false)

	client.RegisterInstance("auth-server", instance)

	go func() {

		for {
			
			err := client.SendHeartbeat(instance.App, instance.HostName)

			if err != nil {
				log.Print("Error: Eureka heartbeat failed ")
			}else {
				log.Print("Info: Eureka heartbeat success")
			}

			time.Sleep(time.Second * 5)

		}

	}()

}