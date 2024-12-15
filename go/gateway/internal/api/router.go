package api

import (
	"fmt"
	"net/http"
	"riki_gateway/internal/services"

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

	return http.ListenAndServe(fmt.Sprintf(":%s", s.svcs.Config().GatewayConfig.GatewayServerPort), s.router)
}
