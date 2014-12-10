package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("workers", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen")
)

func main() {
	//Parse command-line flags
	flag.Parse()
	//Start Dispatcher
	fmt.Println("Starting Dispatcher")
	StartDispatcher(*NWorkers)
	//Register collector
	fmt.Println("Registering Collector")
	http.HandleFunc("/work", Collector)

	//Start Server
	fmt.Println("HTTP Listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
