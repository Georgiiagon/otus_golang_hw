package memorystorage

import (
	"context"
	"sync"
)

type Storage struct {
	// TODO
	mu sync.RWMutex
}

func (s *Storage) Connect(ctx context.Context) error {
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return nil
}

func New() *Storage {
	return &Storage{}
}

// TODO
