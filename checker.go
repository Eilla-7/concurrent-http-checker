package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var printMutex sync.Mutex

func checkURL(url string, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	start := time.Now()
	resp, err := client.Get(url)
	elapsed := time.Since(start)

	printMutex.Lock()
	defer printMutex.Unlock()

	if err != nil {
		fmt.Printf("%s -> ERROR: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%s -> %d (Time: %v)\n", url, resp.StatusCode, elapsed)
}
