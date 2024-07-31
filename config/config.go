package config

import (
	"errors"
	"os"
)

var (
	LoadBalancerAddr = ""
	APIAddr          = ""
	MongoURI         = ""
)

func Load() error {
	LoadBalancerAddr = os.Getenv("LOADBALANCER_ADDR")
	if LoadBalancerAddr == "" {
		return errors.New("LOADBALANCER_ADDR is empty")
	}

	APIAddr = os.Getenv("API_ADDR")
	if APIAddr == "" {
		return errors.New("API_ADDR is empty")
	}

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		return errors.New("MONGO_URI is empty")
	}

	return nil
}
