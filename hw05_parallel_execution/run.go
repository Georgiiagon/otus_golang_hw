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
	ch := make(chan Task, n)
	errCount := int64(m)
	s := settings{
		ch:       ch,
		errCount: &errCount,
		wg:       &sync.WaitGroup{},
	}

	s.wg.Add(len(tasks))

	for i := 0; i < n; i++ {
		go worker(&s)
	}

	for _, task := range tasks {
		if atomic.LoadInt64(s.errCount) >= 0 {
			s.ch <- task
			continue
		}
		close(s.ch)
		return ErrErrorsLimitExceeded
	}
	close(s.ch)
	s.wg.Wait()

	return nil
}

func worker(s *settings) {
	for {
		task, ok := <-s.ch
		if !ok {
			return
		}
		err := task()
		if err != nil {
			atomic.AddInt64(s.errCount, -1)
		}
		s.wg.Done()
	}
}
