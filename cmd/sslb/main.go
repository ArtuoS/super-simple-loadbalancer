package main

import (
	"log"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/api"
)

func main() {
	r := api.NewRouter()
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Panic(err)
	}
}
