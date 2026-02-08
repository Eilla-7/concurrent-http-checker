package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://google.com", "https://github.com", "https://invalid-url"}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, &wg, client)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}
