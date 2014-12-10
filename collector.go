package main

import (
	"fmt"
	"net/http"
	"time"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "Delay must be between 1-10s", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Must Specify a name", http.StatusBadRequest)
		return
	}

	work := WorkRequest{Name: name, Delay: delay}
	WorkQueue <- work
	fmt.Println("Work request queued!")

	w.WriteHeader(http.StatusCreated)
	return
}
