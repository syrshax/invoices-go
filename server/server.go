package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port string
	mux  http.ServeMux
}

// This works like a staticmethod. You create a Server with this function.
func NewServer(port string) *Server {
	return &Server{
		port: port,
		mux:  *http.NewServeMux(),
	}
}

func (s *Server) AddHandler(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%v", s.port)
	log.Printf("Server Starting at port: %v \n", addr)

	return http.ListenAndServe(addr, &s.mux)
}
