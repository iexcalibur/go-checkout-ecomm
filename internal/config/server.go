package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	port   string
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

func (s *Server) Start() error {
	log.Printf("Server starting on port %s...", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.router)
}
