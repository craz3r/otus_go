package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	tasksCh := make(chan Task)
	var mu sync.Mutex
	var wg sync.WaitGroup

	errCount := 0

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range tasksCh {
				if err := task(); err != nil {
					mu.Lock()
					errCount++
					mu.Unlock()
				}
			}
		}()
	}

	for _, task := range tasks {
		mu.Lock()
		exceeded := errCount >= m
		mu.Unlock()
		if exceeded {
			break
		}
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()

	mu.Lock()
	exceeded := errCount >= m
	mu.Unlock()

	if exceeded {
		return ErrErrorsLimitExceeded
	}

	return nil
}
