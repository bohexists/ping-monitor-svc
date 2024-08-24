package workerpool

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Job represents a unit of work (URL)
type Job struct {
	URL string
}

// Result represents the outcome of processing a Job
type Result struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

// Info provides a formatted string with the result information
func (r Result) Info() string {
	if r.Error != nil {
		return fmt.Sprintf("Error - %s - %s", r.URL, r.Error.Error())
	}

	return fmt.Sprintf("Success - %s - Status: %d, Response Time: %s", r.URL, r.StatusCode, r.ResponseTime.String())
}

// Pool represents a pool of workers processing jobs concurrently
type Pool struct {
	worker       *worker
	workersCount int

	jobs    chan Job
	results chan Result

	wg      *sync.WaitGroup
	stopped bool
}

// New creates a Pool with a set number of workers
func New(workersCount int, timeout time.Duration, results chan Result) *Pool {
	return &Pool{
		worker:       newWorker(timeout),
		workersCount: workersCount,
		jobs:         make(chan Job),
		results:      results,
		wg:           new(sync.WaitGroup),
	}
}

// Init starts all the workers in the pool
func (p *Pool) Init() {
	for i := 0; i < p.workersCount; i++ {
		go p.initWorker(i)
	}
}

// Push adds a new job to the pool
func (p *Pool) Push(j Job) {
	if p.stopped {
		return
	}

	p.jobs <- j
	p.wg.Add(1)
}

// Stop gracefully stops all workers by closing the jobs channel
func (p *Pool) Stop() {
	p.stopped = true
	close(p.jobs)
	p.wg.Wait()
}

// initWorker processes jobs from the channel
func (p *Pool) initWorker(id int) {
	for job := range p.jobs {
		time.Sleep(time.Second)
		p.results <- p.worker.process(job)
		p.wg.Done()
	}

	log.Printf("Worker ID %d finished proccesing!", id)
}
