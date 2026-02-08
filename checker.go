package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkURL(url string, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("%s -> ERROR: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	fmt.Printf("%s -> %d (Time: %v)\n", url, resp.StatusCode, elapsed)
}
