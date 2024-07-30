package presenter

import "go.mongodb.org/mongo-driver/bson/primitive"

type Server struct {
	ID        primitive.ObjectID `json:"id"`
	DNS       string             `json:"dns"`
	CallCount int64              `json:"call_count"`
}
