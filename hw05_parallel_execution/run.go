package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type settings struct {
	ch       chan Task
	errCount *int64
	wg       *sync.WaitGroup
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan Task, len(tasks))
	errCount := int64(m)
	s := settings{
		ch:       ch,
		errCount: &errCount,
		wg:       &sync.WaitGroup{},
	}
	s.wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer s.wg.Done()
			worker(&s)
		}()
	}

	for _, t := range tasks {
		s.ch <- t
	}

	close(s.ch)
	s.wg.Wait()

	if atomic.LoadInt64(s.errCount) < 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(s *settings) {
	for t := range s.ch {
		if atomic.LoadInt64(s.errCount) <= 0 {
			return
		}

		err := t()

		if err != nil {
			atomic.AddInt64(s.errCount, -1)
		}
	}
}
