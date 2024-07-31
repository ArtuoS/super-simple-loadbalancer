package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/usecase/balancer"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

func MakeProxyHandler(r *chi.Mux, service balancer.UseCase) {
	r.Use(func(h http.Handler) http.Handler {
		return handleRequest(h, service)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
}

func handleRequest(next http.Handler, service balancer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		r.Body = io.NopCloser(bytes.NewReader(body))

		balancer, err := service.GetFirst(bson.D{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		server := balancer.GetServerWithSmallerCallCounter()
		server.CallCount++
		service.UpdateServer(balancer.ID, server.ID, server.DNS, server.Capacity, server.CallCount)

		url := fmt.Sprintf("%s%s", server.DNS, r.RequestURI)

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
