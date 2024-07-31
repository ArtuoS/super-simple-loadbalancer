package balancer

import (
	"context"
	"errors"
	"time"

	"github.com/ArtuoS/super-simple-loadbalancer/entity"
	"github.com/ArtuoS/super-simple-loadbalancer/infra/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	repo *repository.BalancerMongoDB
}

func NewService(repo *repository.BalancerMongoDB) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PushServer(id primitive.ObjectID, dns string, capacity int) error {
	filter := bson.D{{Key: "_id", Value: id}}
	balancer, err := s.GetFirst(filter)
	if err != nil {
		return errors.New(err.Error())
	}

	if balancer.ExceedMaxCapacity(capacity) {
		return errors.New("balancer capacity was exceeded")
	}

	server := entity.NewServer(dns, 0, capacity)
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "servers", Value: server},
		}},
	}
	filter = bson.D{{Key: "_id", Value: id}}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.repo.PushServer(context, filter, update)
	if err != nil {
		return errors.New(err.Error())
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

	cursor, err := s.repo.Search(context, filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var results []*entity.Balancer
	if err = cursor.All(context, &results); err != nil {
		return nil, errors.New(err.Error())
	}

	return results, nil
}

func (s *Service) Search(filter primitive.D) ([]*entity.Balancer, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.repo.Search(context, filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var results []*entity.Balancer
	if err = cursor.All(context, &results); err != nil {
		return nil, errors.New(err.Error())
	}

	return results, nil
}

func (s *Service) GetFirst(filter primitive.D) (*entity.Balancer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.repo.Search(ctx, filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result entity.Balancer
		if err := cursor.Decode(&result); err != nil {
			return nil, errors.New(err.Error())
		}
		return &result, nil
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.New(err.Error())
	}

	return nil, nil
}

func (s *Service) UpdateServer(balancerId primitive.ObjectID, serverId primitive.ObjectID, dns string, capacity int, callCount int64) error {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "servers.$[elem].dns", Value: dns},
			{Key: "servers.$[elem].capacity", Value: capacity},
			{Key: "servers.$[elem].callcount", Value: callCount},
		}},
	}
	filter := bson.D{{Key: "_id", Value: balancerId}}
	arrayFilters := options.ArrayFilters{
		Filters: bson.A{bson.D{{Key: "elem._id", Value: serverId}}},
	}
	options := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.repo.Update(ctx, filter, update, &options)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
