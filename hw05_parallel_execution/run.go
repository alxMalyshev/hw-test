package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error


func producer(tasksChan chan<- Task, tasks []Task, wg *sync.WaitGroup) {
  for _,task := range tasks {
    go func(){
      defer wg.Done()
      tasksChan <- task
    }()
  }
}

func consumer(n int, m int, tasksChan <-chan Task, wg *sync.WaitGroup) error {
	var errCount int32

	fmt.Println("test")
	for i:=0; i < n; i ++{
			fmt.Println("Error count is:", errCount)
			go func() {
				defer wg.Done()
				for task := range tasksChan {
					if err := task(); err != nil {
						atomic.AddInt32(&errCount,1)
						fmt.Println(errCount)	
					} 
				} 
			}()
	}
	fmt.Println("Result:", errCount)
	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	} else {
		return nil
	}
}

func Run(tasks []Task, n, m int) error {
  wgProducer := sync.WaitGroup{}
  wgConsumer := sync.WaitGroup{}
  tasksChan := make(chan Task)

  wgProducer.Add(len(tasks))
  producer(tasksChan, tasks, &wgProducer)

  wgConsumer.Add(n)
  err := consumer(n, m, tasksChan, &wgConsumer)
  
  wgProducer.Wait()
  close(tasksChan)
  wgConsumer.Wait()

  return err
}