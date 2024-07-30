package main

import (
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"github.com/ArtuoS/super-simple-loadbalancer/usecase/balancer"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	database.InitializeMongoDBClient()

	r := chi.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		service := &balancer.Service{}
		id, _ := primitive.ObjectIDFromHex("66a8cd66c3cde4d3662d46bc")
		service.PushServer(id, "https://www.mozilla.org/")
	})

	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		panic(err)
	}
}
