package main

import (
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/api/handler"
	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"github.com/ArtuoS/super-simple-loadbalancer/infra/repository"
	"github.com/ArtuoS/super-simple-loadbalancer/usecase/balancer"
	"github.com/go-chi/chi"
)

func main() {
	db := database.NewMongoDBClient()
	repo := repository.NewBalancerMongoDB(db)
	service := balancer.NewService(repo)

	r := chi.NewRouter()
	handler.MakeBalancerHandler(r, service)
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		panic(err)
	}
}
