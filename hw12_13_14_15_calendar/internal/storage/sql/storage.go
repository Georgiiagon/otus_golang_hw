package sqlstorage

import (
	"context"
	"database/sql"

	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
	storagemodels "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/models"
)

type Storage struct {
	dsn string
	db  *sql.DB
}

func New(config config.DatabaseConf) *Storage {
	return &Storage{
		db:  &sql.DB{},
		dsn: config.DSN,
	}
}

func (s *Storage) Connect(ctx context.Context) (err error) {
	s.db, err = sql.Open("pgx", s.dsn)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) CreateEvent(event storagemodels.Event) (*storagemodels.Event, error) {
	return &event, nil
}

func (s *Storage) UpdateEvent(event storagemodels.Event) (*storagemodels.Event, error) {
	return &event, nil
}

func (s *Storage) DeleteEvent(id int) error {
	return nil
}

func (s *Storage) GetEvents() []*storagemodels.Event {
	eventsSlice := make([]*storagemodels.Event, 0)

	return eventsSlice
}
