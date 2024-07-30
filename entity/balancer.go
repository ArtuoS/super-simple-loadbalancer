package entity

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Balancer struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Servers []*Server          `json:"servers"`
}

func NewBalancer() *Balancer {
	return &Balancer{
		ID: primitive.NewObjectID(),
	}
}

func (b *Balancer) PushServer(server *Server) {
	b.Servers = append(b.Servers, server)
}

func (b *Balancer) getServerWithSmallerCallCounter() *Server {
	if len(b.Servers) == 1 {
		return b.Servers[0]
	}

	s := b.Servers[0]
	for _, v := range b.Servers {
		if v.CallCount < s.CallCount {
			s = v
		}
	}

	log.Printf("Servidor que será utilizado: %s", s.ToString())
	return s
}

func (b *Balancer) HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		r.Body = io.NopCloser(bytes.NewReader(body))

		s := b.getServerWithSmallerCallCounter()
		s.CallCount++
		url := fmt.Sprintf("%s%s", s.DNS, r.RequestURI)

		log.Printf("Requisição sendo redirecionada para: %s", url)

		proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		proxyReq.Header = make(http.Header)
		for h, val := range r.Header {
			proxyReq.Header[h] = val
		}

		resp, err := http.DefaultClient.Do(proxyReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for h, vals := range resp.Header {
			for _, val := range vals {
				w.Header().Add(h, val)
			}
		}

		w.WriteHeader(resp.StatusCode)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
