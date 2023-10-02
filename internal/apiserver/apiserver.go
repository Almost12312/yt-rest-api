package apiserver

import "log/slog"

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

func (s *APIServer) Start() error {
	return nil
}
