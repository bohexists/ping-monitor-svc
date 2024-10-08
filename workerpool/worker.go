package workerpool

import (
	"net/http"
	"time"
)

// worker makes HTTP requests
type worker struct {
	client *http.Client
}

// newWorker creates a worker with a timeout
func newWorker(timeout time.Duration) *worker {
	return &worker{
		&http.Client{
			Timeout: timeout,
		},
	}
}

// process performs the job (HTTP GET)
func (w worker) process(j Job) Result {
	result := Result{URL: j.URL}

	now := time.Now()

	resp, err := w.client.Get(j.URL)
	if err != nil {
		result.Error = err
		return result
	}

	result.StatusCode = resp.StatusCode
	result.ResponseTime = time.Since(now)

	return result
}
