package presenter

import "go.mongodb.org/mongo-driver/bson/primitive"

type Balancer struct {
	ID      primitive.ObjectID `json:"id"`
	Servers []*Server          `json:"servers"`
}
