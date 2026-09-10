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

	tasksCh := make(chan Task, len(tasks))
	errorsCh := make(chan error, len(tasks))
	stopCh := make(chan struct{})

	wg := sync.WaitGroup{}

	go func() {
		for _, task := range tasks {
			tasksCh <- task
		}
		close(tasksCh)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-stopCh:
					return
				case task, ok := <-tasksCh:
					if !ok {
						return
					}
					err := task()
					if err != nil {
						select {
						case <-stopCh:
							return
						case errorsCh <- err:
						}
					}
				}
			}
		}()
	}

	go func() {

		close(errorsCh)
	}()

	errorsCount := 0

	for range errorsCh {
		errorsCount++

		if errorsCount >= m {
			close(stopCh)
			break
		}
	}

	wg.Wait()

	if errorsCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
