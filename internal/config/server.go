package config

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port    string
	router  *mux.Router
	handler http.Handler
}

func NewServer(port string) *Server {
	return &Server{
		router: mux.NewRouter(),
		port:   port,
	}
}

func (s *Server) Router() *mux.Router {
	return s.router
}

func (s *Server) SetHandler(h http.Handler) {
	s.handler = h
}

func (s *Server) Start() error {
	if s.handler == nil {
		s.handler = s.router
	}
	return http.ListenAndServe(":"+s.port, s.handler)
}
