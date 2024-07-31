package balancer

import (
	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Get(id primitive.ObjectID) ([]*entity.Balancer, error)
	Search(filter primitive.D) ([]*entity.Balancer, error)
	GetFirst(filter primitive.D) (*entity.Balancer, error)

	PushServer(id primitive.ObjectID, dns string, capacity int) error
	UpdateServer(balancerId primitive.ObjectID, serverId primitive.ObjectID, dns string, capacity int, callCount int64) error
}
