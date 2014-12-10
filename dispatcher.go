package main

import "fmt"

var WorkerQueue chan chan WorkRequest

//StartDispatcher starts the dispatch goroutine
func StartDispatcher(nworkers int) {
	//Initialize the channel we are put the workers work channels into
	WorkerQueue = make(chan chan WorkRequest, nworkers)
	//Create all workers
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	//Actual Dispatching
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
