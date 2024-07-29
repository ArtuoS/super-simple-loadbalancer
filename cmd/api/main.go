package main

import (
	"context"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"github.com/go-chi/chi"
)

func main() {
	database.InitializeMongoDBClient()

	r := chi.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		balancer := entity.NewBalancer()
		server1 := entity.NewServer("https://www.google.com", 0)
		server2 := entity.NewServer("https://duckduckgo.com", 0)

		balancer.PushServer(server1)
		balancer.PushServer(server2)

		_, err := database.Client.Database("loadbalancer").Collection(database.Balancers).InsertOne(context.TODO(), balancer)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", r)
}
