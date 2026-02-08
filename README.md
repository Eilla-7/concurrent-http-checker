# Concurrent HTTP URL Checker (Go)

This project is a simple Go application that checks multiple URLs concurrently using **goroutines**.  
---

##  Project Overview

The application:
- Takes a list of URLs
- Sends HTTP GET requests to all of them concurrently
- Measures response time for each request
- Uses a shared `http.Client` with a timeout
- Waits for all goroutines to finish before exiting

---

## How It Works

1. A list of URLs is defined in `main`
2. For each URL:
   - A goroutine is started
   - An HTTP request is sent
   - The response time is measured
3. A `sync.WaitGroup` is used to ensure the program waits for all goroutines to complete
4. Results are printed as soon as each request finishes

> Output order may vary due to concurrent execution.

---

## Code Highlights

- **Goroutines** for concurrency
- **sync.WaitGroup** for synchronization
- **http.Client with timeout** to avoid hanging requests
- Response time measurement using `time.Since`

---
