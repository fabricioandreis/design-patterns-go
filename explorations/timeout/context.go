package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type response struct {
	url     string
	status  string
	latency time.Duration
}

func get(c context.Context, url string, ch chan<- response) {
	start := time.Now()
	req, err := http.NewRequestWithContext(c, "GET", url, nil)
	if err != nil {
		log.Fatalf("Error when creating request with context: %s", err.Error())
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error when getting url "+url, err)
	}
	ch <- response{url: url, status: resp.Status, latency: time.Since(start).Round(time.Millisecond)}
}

func main() {

	urls := []string{
		"https://google.com",
		"https://microsoft.com",
		"https://facebook.com",
		"https://linkedin.com",
		"http://localhost:5000",
	}

	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ch := make(chan response, len(urls))

	for _, u := range urls {
		go get(c, u, ch)
	}

	for range urls {
		resp := <-ch
		log.Printf("GET %s: %s in %s\n", resp.url, resp.status, resp.latency)
	}
}
