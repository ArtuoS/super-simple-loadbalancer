package domain

import (
	"fmt"
)

type Server struct {
	ID        int
	URL       string
	callCount int
}

func NewServer(id int, url string) *Server {
	return &Server{
		ID:  id,
		URL: url,
	}
}

func (s *Server) Call() {
	fmt.Printf("Calling %d %s\n", s.ID, s.URL)
	s.IncreaseCallCounter()
}

func (s *Server) IncreaseCallCounter() {
	s.callCount++
}

func (s *Server) CallCount() int {
	return s.callCount
}

func (s *Server) ToString() string {
	return fmt.Sprintf("ID: %d, URL: %s, Call Count: %d\n", s.ID, s.URL, s.CallCount())
}
