package entity

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DNS       string             `json:"dns"`
	CallCount int64              `json:"call_count"`
}

func NewServer(url string, callCount int64) *Server {
	return &Server{
		ID:        primitive.NewObjectID(),
		DNS:       url,
		CallCount: callCount,
	}
}

func (s *Server) ToString() string {
	return fmt.Sprintf("URL: %s, Call Count: %d\n", s.DNS, s.CallCount)
}
