package apiserver

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type APIServer struct {
	config *Config
	logger slog.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: slog.Logger{},
	}
}

func (s *APIServer) Start(r *mux.Router) error {
	return http.ListenAndServe(s.config.BindAddr, r)
}
