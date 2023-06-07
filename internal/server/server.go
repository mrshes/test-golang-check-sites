package server

import (
	"net/http"
	"test-golang-check-sites/internal/models"
	"test-golang-check-sites/internal/services"
)

type Server struct {
	Services   *services.Services
	Statistics *models.Statistics
}

func NewServer(services *services.Services) *Server {
	return &Server{
		Services:   services,
		Statistics: models.NewStatistics(),
	}
}

func (s *Server) Run(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.index)
	mux.HandleFunc("/min", s.getMin)
	mux.HandleFunc("/max", s.getMax)

	mux.HandleFunc("/admin", s.getStatistics)

	return http.ListenAndServe(":"+port, mux)
}
