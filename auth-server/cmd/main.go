package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	log.Print("Auth server started")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	registerWithEureka()

	http.ListenAndServe(":8083", r)
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