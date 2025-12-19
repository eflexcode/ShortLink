package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	log.Print("User server started")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	

	registerWithEureka()

	http.ListenAndServe(":8082", r)

}

func registerWithEureka() {

	client := eureka.NewClient([]string{"http://localhost:8081/eureka"})

	instance := eureka.NewInstanceInfo("localhost:8082", "user-server", "127.0.0.1", 8082, 30, false)

	client.RegisterInstance("user-server", instance)

	go func() {

		for {
			err := client.SendHeartbeat(instance.App, instance.HostName)

			if err != nil {
				log.Print("Error: Eureka heartbeat failed "+err.Error())
			}else {
				log.Print("Info: Eureka heartbeat success")
			}

			time.Sleep(time.Second * 5)

		}

	}()

}
