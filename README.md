# Ping Monitor Service

The Ping Monitor Service is a simple tool that continuously monitors the availability and response times of specified URLs. It utilizes a worker pool to efficiently handle multiple concurrent requests.

## Features

- Periodically checks the availability of URLs.
- Logs the status code, response time, and any errors encountered during the request.
- Configurable worker pool for handling multiple URLs concurrently.
- Graceful shutdown on receiving termination signals.

## Project Structure

- **main.go**: The entry point of the service. It initializes the worker pool and manages the job generation and result processing.
- **workerpool/pool.go**: Contains the `Pool` struct that manages the distribution of jobs to workers and handles the lifecycle of the worker pool.
- **workerpool/worker.go**: Contains the `worker` struct responsible for making HTTP requests and collecting the results.

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/bohexists/ping-monitor-svc.git
   cd ping-monitor-svc
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Usage

To run the Ping Monitor Service, execute the following command:

```bash
go run main.go
```

The service will start monitoring the specified URLs every minute and print the results to the console.

### Configuration

You can configure the following constants in `main.go`:

- `INTERVAL`: The interval between each set of checks (default: 1 minute).
- `REQUEST_TIMEOUT`: The timeout for each HTTP request (default: 1 second).
- `WORKERS_COUNT`: The number of workers handling the requests (default: 2 workers).

### Example Output

```
Success - https://google.com/ - Status: 200, Response Time: 123ms
Success - https://golang.org/ - Status: 200, Response Time: 145ms
```

### Graceful Shutdown

The service listens for termination signals (`SIGTERM`, `SIGINT`). Upon receiving such a signal, it will stop accepting new jobs and wait for all ongoing jobs to complete before shutting down.


### Contributions

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.

### Contact

For any inquiries, please contact https://github.com/bohexists.