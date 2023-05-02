package service

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Service struct {
	config *Config
	logger *logrus.Logger
	router *chi.Mux
}

func New(config *Config) *Service {
	return &Service{
		config: config,
		logger: logrus.New(),
		router: chi.NewRouter(),
	}
}

func (s *Service) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting user service")

	return http.ListenAndServe(s.config.BindAddr, s.router) //nolint:gosec  // tests start
}

func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Service) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Service) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello") //nolint:errcheck  // tests handle
	}
}
