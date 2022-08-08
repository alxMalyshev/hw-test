package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

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

	fmt.Println("Channel len is:", len(tasksChan))

	wg.Add(n)
	for i := 0; i < n; i++ {
		fmt.Println("ErrorCount is:", errCount)
		go func() {
			defer wg.Done()
			for task := range tasksChan {
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}
	fmt.Println(atomic.LoadInt32(&errCount))
	return atomic.LoadInt32(&errCount)
}

func Run(tasks []Task, n, m int) error {
	errCount := consumer(n, m, producer(tasks))
	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	} else {
		return nil
	}
}
