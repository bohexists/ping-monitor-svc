# Ping Monitor Service (ping-monitor-svc)

Ping Monitor Service is a Go-based application designed to monitor the availability and response time of specified URLs. The service uses a worker pool to perform HTTP requests periodically and can send notifications to a specified Telegram chat.

## Features

- **URL Monitoring**: Periodically checks the availability and response time of configured URLs.
- **Worker Pool**: Efficiently manages multiple concurrent HTTP requests using a configurable number of workers.
- **Result Logging**: Logs the status, response time, and errors (if any) for each monitored URL.
- **Telegram Notifications**: Sends notifications to a specified Telegram chat when a URL is checked, with details of the response or error.
- **Graceful Shutdown**: Handles OS signals for a clean shutdown of the service.

## Installation

To install the Ping Monitor Service, clone the repository and ensure you have Go installed on your machine.

```sh
git clone https://github.com/bohexists/ping-monitor-svc.git
cd ping-monitor-svc
go mod tidy
```

## Configuration

### Environment Variables

Before running the service, set up the following environment variables:

- `BOT_TOKEN`: Your Telegram bot token.
- `CHAT_ID`: The chat ID where notifications will be sent.

You can set these variables in your shell or include them in your code if you prefer hardcoding (not recommended for production).

### URLs to Monitor

The URLs to monitor are currently hardcoded in the `main.go` file:

```go
var urls = []string{
    "https://google.com/",
    "https://calendar.google.com/",
    "https://mail.google.com/",
    "https://drive.google.com/",
    "https://translate.google.co.uk/",
    "https://golang.org/",
    "https://github.com/",
}
```

You can modify this list to include the URLs you need to monitor.

## Usage

To run the service, use the following command:

```sh
go run main.go
```

### Example

When you run the service, it will periodically check the configured URLs and log the results. If you have set up the Telegram bot and chat ID correctly, it will also send the results to your specified Telegram chat.

```sh
2024/08/25 14:35:01 Success - https://google.com/ - Status: 200, Response Time: 112ms
2024/08/25 14:35:01 Success - https://calendar.google.com/ - Status: 200, Response Time: 134ms
2024/08/25 14:35:01 Error - https://mail.google.com/ - Get "https://mail.google.com/": dial tcp: lookup mail.google.com: no such host
```

The same messages will be sent to your Telegram chat.

## Code Structure

- **`main.go`**: Entry point of the application. Initializes the worker pool, starts generating jobs, processes results, and sends Telegram notifications.
- **`workerpool/`**: Contains the worker pool implementation, including job processing logic.
   - `pool.go`: Defines the `Pool` struct and its methods.
   - `worker.go`: Implements the worker that processes each job.
- **`telegram/`**: Contains the Telegram notification logic.
   - `telegram.go`: Defines the `Telegram` struct and methods for sending notifications.

## Telegram Integration

The Telegram integration allows you to receive real-time notifications in your Telegram chat. You need to configure your bot token and chat ID:

```go
botToken := "YOUR_BOT_TOKEN"
chatID := int64(YOUR_CHAT_ID)
telegramSender, err := telegram.NewSender(botToken, chatID)
if err != nil {
    log.Fatal(err)
}
```

In the `proccessResults` function, the results are sent to the Telegram chat:

```go
err := telegramSender.SendNotification(result.Info())
if err != nil {
    log.Println("Failed to send notification:", err)
}
```

## Future Enhancements

- **Dashboard Integration**: Adding a web-based dashboard for real-time monitoring.
- **Advanced Error Handling**: Implementing more granular error tracking and alerting.
- **Extended Notification Options**: Adding support for email and other messaging platforms.
- **Customizable Job Schedules**: Allowing more flexible job scheduling (e.g., CRON-based).
