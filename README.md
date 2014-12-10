#WorkerPool Example

This is an example of a workerpool in Go.

##Summary
A worker pool has three main components:
* **Collector**
  Receives work requests and sends them to the queue (http server)
* **Dispatcher**
  Divies out work request to workers as they become available. In this example, the dispatcher is a goroutine.
  The Dispatcher Goroutine then spawns new go routines to
  send work requests to the queue.  This way, if even if the queue is full, it will not block the collector.
* **Worker**
  Unit for work.  This is where stuff actually gets done.

##Usage
Run the program via executable.  This way it can be done using flags.
```
./workpool [-workers=number of workers] [-http=http address]
```
The default flags will start the program with 4 workers and at localhost:8000.

##Contributors
* **Gian Biondi** <gian@namely.com>
