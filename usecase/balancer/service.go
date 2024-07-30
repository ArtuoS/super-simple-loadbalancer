package balancer

import (
	"context"
	"errors"
	"time"

	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"github.com/ArtuoS/super-simple-loadbalancer/infra/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo *repository.BalancerMongoDB
}

func NewService(repo *repository.BalancerMongoDB) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PushServer(id primitive.ObjectID, dns string) error {
	server := entity.NewServer(dns, 0)
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "servers", Value: server},
		}},
	}
	filter := bson.D{{Key: "_id", Value: id}}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.repo.PushServer(context, filter, update)
	if err != nil {
		return errors.New("failed to push server")
	}

	if result.MatchedCount == 0 {
		return errors.New("no document found with id")
	}

	return nil
}

func (s *Service) Get(id primitive.ObjectID) ([]*entity.Balancer, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.repo.GetServers(context, filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var results []*entity.Balancer
	if err = cursor.All(context, &results); err != nil {
		return nil, errors.New(err.Error())
	}

	return results, nil
}
