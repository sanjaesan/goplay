package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// Job ---
type Job struct {
	Name     string
	Duration time.Duration
}

// Worker ---
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quit       chan bool
}

// NewWorker ---
func NewWorker(workerID int) *Worker {
	workerPool := make(chan chan Job)
	return &Worker{
		id:         workerID,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quit:       make(chan bool),
	}
}

// Start ---
func (s Worker) Start() {
	go func() {
		for {
			// add job to the worker pool
			s.workerPool <- s.jobQueue
			select {
			case job := <-s.jobQueue:
				fmt.Printf("worker%d: completed:%s!\n", s.id, job)
				time.Sleep(job.Duration)

			case <-s.quit:
				fmt.Printf("Worker%d stopping...", s.id)
				return
			}
		}
	}()
}

// Stop ---
func (s Worker) Stop() {
	go func() {
		s.quit <- true
	}()
}

// Dispatcher ---
type Dispatcher struct {
	workerPool chan chan Job
	maxWorker  int
	jobQueue   chan Job
}

//NewDispatcher ---
func NewDispatcher(maxWorker int) *Dispatcher {
	jobQueue := make(chan Job)
	workerPool := make(chan chan Job)
	return &Dispatcher{
		workerPool: workerPool,
		maxWorker:  maxWorker,
		jobQueue:   jobQueue,
	}
}

// Run ---
func (s Dispatcher) Run() {
	for i := 0; i < s.maxWorker; i++ {
		worker := NewWorker(i + 1)
		worker.Start()
	}
	go s.Dispatch()
}

// Dispatch ---
func (s Dispatcher) Dispatch() {
	for {
		select {
		case job := <-s.jobQueue:
			go func() {
				fmt.Printf("Fetching Worker Pool for %s\n:", job.Name)
				pool := <-s.workerPool
				fmt.Printf("Adding %s to worker pool\n", job.Name)
				pool <- job
			}()
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value"+err.Error(), http.StatusBadRequest)
		return
	}
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "Delay vaue must be between 1 and 10 seconds", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "empty name, you must specify a name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Duration: delay}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	// var (
	// 	maxWorker    = flag.Int("max_worker", 5, "Maximum nuber of workers")
	// 	maxQueueSize = flag.Int("max_queue_size", 200, "maximum queue size")
	// 	port         = flag.String("Port number", "3000", "The server port")
	// )
	flag.Parse()

}
