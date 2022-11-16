package server

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	Port string
}

type JsonMessage struct {
	Message string `json:"message"`
}

func New(p string) *Server {
	return &Server{Port: p}
}

func (s *Server) GetMessage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		jm := JsonMessage{"Hello World"}

		response, _ := json.Marshal(jm)
		w.Write(response)
	}
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		jm := JsonMessage{"Server is work"}

		response, _ := json.Marshal(jm)
		w.Write(response)
	}
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Page not found"))
}