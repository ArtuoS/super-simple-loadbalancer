package api

import (
	"github.com/ArtuoS/super-simple-loadbalancer/internal/handlers"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(handlers.HandleRequest)
	return r
}
