package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArtuoS/super-simple-loadbalancer/api/presenter"
	"github.com/ArtuoS/super-simple-loadbalancer/usecase/balancer"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeBalancerHandler(r *chi.Mux, service balancer.UseCase) {
	r.Put("/api/v1/balancer/{id}/server", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			DNS      string `json:"dns"`
			Capacity int    `json:"capacity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Panic(err)
		}
		id := chi.URLParam(r, "id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Panic(err)
		}
		if err := service.PushServer(oid, input.DNS, input.Capacity); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})

	r.Get("/api/v1/balancer/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Panic(err)
		}
		results, err := service.Get(oid)
		if err != nil {
			log.Panic(err)
		}

		var balancers []*presenter.Balancer
		for _, b := range results {
			balancer := &presenter.Balancer{
				ID: b.ID,
			}

			for _, s := range b.Servers {
				balancer.Servers = append(balancer.Servers, &presenter.Server{
					ID:        s.ID,
					DNS:       s.DNS,
					CallCount: s.CallCount,
					Capacity:  s.Capacity,
				})
			}

			balancers = append(balancers, balancer)
		}

		if err := json.NewEncoder(w).Encode(balancers); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
