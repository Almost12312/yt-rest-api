package apiserver

import (
	"github.com/gorilla/mux"
	"log/slog"
	"web-server-yt/internal/storage"
)

type APIServer struct {
	config  *Config
	logger  slog.Logger
	storage *storage.Storage
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: slog.Logger{},
	}
}

func (s *APIServer) Start(r *mux.Router) error {
	err := s.configureStore()
	if err != nil {
		return err
	}

	return nil
	//return http.ListenAndServe(s.config.BindAddr, r)
}

func (s *APIServer) configureStore() error {
	const op = "apiserver.apiserver.configureStore"

	st := storage.New(s.config.Storage)

	if err := st.Open(); err != nil {
		//s.logger.Error("error in open store in op:")
		return err
	}

	s.storage = st

	//s.logger.Info("Storage success created! in op:")

	return nil
}
