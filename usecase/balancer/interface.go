package balancer

import (
	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Get(id primitive.ObjectID) ([]*entity.Balancer, error)
	PushServer(id primitive.ObjectID, dns string, capacity int) error
}
