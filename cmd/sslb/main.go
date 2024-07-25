package main

import (
	"fmt"
	"time"

	"github.com/ArtuoS/super-simple-loadbalancer/internal/domain"
)

func main() {
	balancer := domain.NewBalancer()
	server1 := domain.NewServer(1, "https://www.google.com")
	server2 := domain.NewServer(2, "https://www.youtube.com")
	server3 := domain.NewServer(3, "https://www.facebook.com")

	balancer.PushServer(server1)
	balancer.PushServer(server2)
	balancer.PushServer(server3)

	start := time.Now()
	for i := 0; i < 100000; i++ {
		balancer.HandleRequest()
	}
	total := time.Since(start)

	for _, v := range balancer.Servers {
		fmt.Println(v.ToString())
	}

	fmt.Println(total)
}
