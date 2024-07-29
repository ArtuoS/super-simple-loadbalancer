package api

import (
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	balancer := entity.NewBalancer()
	server1 := entity.NewServer("https://www.google.com", 0)
	server2 := entity.NewServer("https://duckduckgo.com", 0)

	balancer.PushServer(server1)
	balancer.PushServer(server2)

	r := chi.NewRouter()
	r.Use(balancer.HandleRequest)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	return r
}
