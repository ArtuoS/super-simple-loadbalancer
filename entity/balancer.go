package entity

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Balancer struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Servers []*Server          `json:"servers"`
}

func NewBalancer() *Balancer {
	return &Balancer{
		ID: primitive.NewObjectID(),
	}
}

func (b *Balancer) PushServer(server *Server) {
	b.Servers = append(b.Servers, server)
}

func (b *Balancer) GetServerWithSmallerCallCounter() *Server {
	if len(b.Servers) == 1 {
		return b.Servers[0]
	}

	s := b.Servers[0]
	for _, v := range b.Servers {
		if v.CallCount < s.CallCount {
			s = v
		}
	}

	log.Printf("Servidor que será utilizado: %s", s.ToString())
	return s
}

func (b *Balancer) ExceedMaxCapacity(capacity int) bool {
	return (b.ActualCapacity() + capacity) > 100
}

func (b *Balancer) ActualCapacity() int {
	if b.Servers == nil || len(b.Servers) == 0 {
		return 0
	}

	capacity := 0
	for _, s := range b.Servers {
		capacity += s.Capacity
	}
	return capacity
}
