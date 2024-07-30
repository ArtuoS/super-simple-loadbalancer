package entity

import (
	"fmt"
)

type Server struct {
	DNS       string `json:"dns"`
	CallCount int64  `json:"call_count"`
}

func NewServer(url string, callCount int64) *Server {
	return &Server{
		DNS:       url,
		CallCount: callCount,
	}
}

func (s *Server) ToString() string {
	return fmt.Sprintf("URL: %s, Call Count: %d\n", s.DNS, s.CallCount)
}
