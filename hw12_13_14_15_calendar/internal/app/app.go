package app

import (
	"context"

	storagemodels "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/models"
)

type App struct {
	Logger  Logger
	Storage Storage
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
}

type Storage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	CreateEvent(event storagemodels.Event) (*storagemodels.Event, error)
	UpdateEvent(event storagemodels.Event) (*storagemodels.Event, error)
	DeleteEvent(id int) error
	GetEvents() []*storagemodels.Event
}

func New(logger Logger, storage Storage) *App {
	return &App{Logger: logger, Storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, event storagemodels.Event) (*storagemodels.Event, error) {
	return a.Storage.CreateEvent(event)
}

func (a *App) GetEvents(ctx context.Context) []*storagemodels.Event {
	return a.Storage.GetEvents()
}
