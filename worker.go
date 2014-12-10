package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

//NewWorker initializes a new worker
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
	return worker
}

//Start the worker by starting goroutine that is infinite for-select
func (w Worker) Start() {
	go func() {
		for {
			//Worker adds itself to the queue when ready
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				//Receive a work request
				fmt.Printf("Worker%d: Received work, delaying for %f seconds\n", w.ID, work.Delay.Seconds())
				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Howdy, %s!\n", w.ID, work.Name)
			case <-w.QuitChan:
				//Expected Stop
				fmt.Printf("Worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

//Stop tells the worker to stop listening for requests (after it finishes work)
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
