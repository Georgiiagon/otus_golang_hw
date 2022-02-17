package memorystorage

import (
	"context"
	"errors"
	"sync"

	storagemodels "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/models"
)

type Storage struct {
	// TODO
	mu     sync.RWMutex
	LastID int
	Events map[int]*storagemodels.Event
}

var ErrEventNotFound = errors.New("event not found")

func (s *Storage) Connect(ctx context.Context) error {
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return nil
}

func New() *Storage {
	return &Storage{
		Events: make(map[int]*storagemodels.Event),
		LastID: 1,
	}
}

func (s *Storage) CreateEvent(event storagemodels.Event) (*storagemodels.Event, error) {
	s.mu.Lock()
	event.ID = s.LastID
	s.LastID++
	s.Events[event.ID] = &event
	s.mu.Unlock()

	return &event, nil
}

func (s *Storage) UpdateEvent(event storagemodels.Event) (*storagemodels.Event, error) {
	s.mu.RLock()
	if _, ok := s.Events[event.ID]; !ok {
		return nil, ErrEventNotFound
	}
	s.mu.RUnlock()

	s.Events[event.ID] = &event

	return &event, nil
}

func (s *Storage) DeleteEvent(id int) error {
	s.mu.Lock()
	delete(s.Events, id)
	s.mu.Unlock()

	return nil
}

func (s *Storage) GetEvents() []*storagemodels.Event {
	eventsSlice := make([]*storagemodels.Event, 0, len(s.Events))

	s.mu.RLock()
	for _, event := range s.Events {
		eventsSlice = append(eventsSlice, event)
	}
	s.mu.RUnlock()

	return eventsSlice
}
