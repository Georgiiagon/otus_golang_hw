package internalhttp

import (
	"context"
	"net/http"

	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
	storagemodels "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/models"
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
	CreateEvent(ctx context.Context, event storagemodels.Event) (*storagemodels.Event, error)
	GetEvents(ctx context.Context) []*storagemodels.Event
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
	srv := http.Server{Addr: s.Config.App.Host + ":" + s.Config.App.Port, Handler: mux}

	ch := make(chan error)
	go func() {
		err := srv.ListenAndServe()
		ch <- err
	}()

	select {
	case <-ctx.Done():
		s.Logger.Info("Closed by context")
		err := srv.Shutdown(ctx)
		if err != nil {
			return err
		}
	case err := <-ch:
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Info("Stop http server")
	return nil
}

// TODO
