package main

import (
	"log"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/api/handler"
	"github.com/ArtuoS/super-simple-loadbalancer/config"
	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"github.com/ArtuoS/super-simple-loadbalancer/infra/repository"
	"github.com/ArtuoS/super-simple-loadbalancer/usecase/balancer"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	db := database.NewMongoDBClient()
	repo := repository.NewBalancerMongoDB(db)
	service := balancer.NewService(repo)

	r := chi.NewRouter()
	handler.MakeProxyHandler(r, service)
	err = http.ListenAndServe(config.LoadBalancerAddr, r)
	if err != nil {
		log.Panic(err)
	}
}
