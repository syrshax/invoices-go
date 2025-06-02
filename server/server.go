package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port string
	Mux  http.ServeMux
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		Mux:  *http.NewServeMux(),
	}
}

func (s *Server) AddHandler(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	s.Mux.HandleFunc(pattern, handler)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%v", s.port)
	log.Printf("Server Starting at port: %v \n", addr)

	return http.ListenAndServe(addr, &s.Mux)
}
