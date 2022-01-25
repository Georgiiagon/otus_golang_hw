package internalhttp

import (
	"context"
	"net/http"

	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
)

type Server struct {
	Logger Logger
	Config config.Config
	App    Application
}

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type Application interface {
	CreateEvent(context.Context, string, string) error
	GetEvents(ctx context.Context) error
}

func NewServer(logger Logger, app Application, conf config.Config) *Server {
	return &Server{
		Logger: logger,
		Config: conf,
		App:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	handler := apiHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", s.loggingMiddleware(handler.Hello))

	http.ListenAndServe(":3000", mux)
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

// TODO
