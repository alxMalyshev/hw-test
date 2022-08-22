package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkerCount  = errors.New("ivalid worker count")
)

type Task func() error

func producer(tasks []Task) <-chan Task {
	tasksChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	return tasksChan
}

func consumer(n int, m int, tasksChan <-chan Task) int32 {
	var errCount int32
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range tasksChan {
				if err := task(); err != nil {
					if atomic.LoadInt32(&errCount) >= int32(m) {
						break
					}
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}

	wg.Wait()
	return errCount
}

func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidWorkerCount
	}

	if m <= 0 {
		m = len(tasks) + 1
	}

	errCount := consumer(n, m, producer(tasks))

	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
