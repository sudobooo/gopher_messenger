package user

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/sudobooo/gopher_messenger/user_service/internal/database"

	"github.com/sudobooo/gopher_messenger/user_service/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Service struct {
	config   *config.Config
	logger   *logrus.Logger
	router   *chi.Mux
	database *database.Database
}

func New(cfg *config.Config) *Service {
	return &Service{
		config:   cfg,
		logger:   logrus.New(),
		router:   chi.NewRouter(),
		database: database.New(),
	}
}

func (s *Service) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureDatabase(); err != nil {
		return err
	}
	defer db.Close()

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

func (s *Service) configureDatabase() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.database = db

	return nil
}

func (s *Service) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello") //nolint:errcheck  // tests handle
	}
}
