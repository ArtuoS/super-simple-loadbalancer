package main

import (
	"log"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/api/handler"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	handler.MakeProxyHandler(r)
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Panic(err)
	}
}
