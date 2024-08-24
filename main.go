package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bohexists/ping-monitor-svc/telegram"
	"github.com/bohexists/ping-monitor-svc/workerpool"
)

const (
	INTERVAL        = time.Second * 30       // Time between job generations
	REQUEST_TIMEOUT = time.Millisecond * 200 // Timeout for HTTP requests
	WORKERS_COUNT   = 3                      // Number of workers
)

// URLs to monitor
var urls = []string{
	"https://google.com/",
	"https://calendar.google.com/",
	"https://mail.google.com/",
	"https://drive.google.com/",
	"https://translate.google.co.uk/",
	"https://golang.org/",
	"https://github.com/",
}

func main() {
	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WORKERS_COUNT, REQUEST_TIMEOUT, results) // Initialize pool

	workerPool.Init() // Start workers

	botToken := "YOUR_BOT_TOKEN" // Add telegram TOKEN
	chatID := int64(1)           // Add telegram chat ID

	telegramSender, err := telegram.NewSender(botToken, chatID)
	if err != nil {
		log.Fatal(err)
	}

	go generateJobs(workerPool)
	go proccessResults(results, telegramSender)

	quit := make(chan os.Signal, 1) // Handle OS signals
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	workerPool.Stop() // Stop the worker pool
}

// proccessResults prints results from the channel
func proccessResults(results chan workerpool.Result, telegram *telegram.Telegram) {
	go func() {
		for result := range results {
			fmt.Println(result.Info())

			err := telegram.SendNotification(result.Info())
			if err != nil {
				log.Println("Failed to send notification:", err)
			}
		}
	}()
}

// generateJobs periodically adds jobs to the pool
func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}

		time.Sleep(INTERVAL)
	}
}
