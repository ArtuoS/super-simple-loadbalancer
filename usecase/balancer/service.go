package balancer

import (
	"context"
	"errors"
	"time"

	"github.com/ArtuoS/super-simple-loadbalancer/database"
	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
}

func (s *Service) PushServer(id primitive.ObjectID, dns string) error {
	server := entity.NewServer(dns, 0)
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$push", bson.D{
			{"servers", server},
		}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := database.Client.Database("loadbalancer").Collection(database.Balancers).UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to push server")
	}

	if result.MatchedCount == 0 {
		return errors.New("no document found with id")
	}

	return nil
}
