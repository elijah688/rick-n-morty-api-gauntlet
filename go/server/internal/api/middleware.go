package api

import (
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Middleware() {
	s.router.Use(middleware.Logger)
	s.router.Use(timeoutMiddleware)
	s.router.Use(corseMiddleware)
}
