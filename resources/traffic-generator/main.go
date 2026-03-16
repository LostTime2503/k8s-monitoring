package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	urlPtr := flag.String("url", "http://localhost:8080", "Target URL to generate traffic for")
	flag.Parse()

	targetURL := *urlPtr
	fmt.Printf("Starting traffic generator targeting: %s\n", targetURL)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for {
		resp, err := client.Get(targetURL)
		if err != nil {
			fmt.Printf("Error requesting %s: %v\n", targetURL, err)
		} else {
			fmt.Printf("Request sent to %s. Status: %s\n", targetURL, resp.Status)
			resp.Body.Close()
		}

		// Sleep for a random duration between 100ms and 2000ms
		sleepDuration := time.Duration(rand.Intn(1900)+100) * time.Millisecond
		time.Sleep(sleepDuration)
	}
}