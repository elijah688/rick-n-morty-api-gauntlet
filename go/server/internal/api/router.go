package api

import (
	"net/http"
	"riki/internal/services"

	"github.com/go-chi/chi"
)

type Server struct {
	router *chi.Mux
	svcs   *services.Services
}

func New(svcs *services.Services) *Server {
	return &Server{chi.NewRouter(), svcs}
}

func (s *Server) Start() error {
	s.Middleware()
	s.Routes()

	// to do: get from config
	return http.ListenAndServe(":8080", s.router)
}
