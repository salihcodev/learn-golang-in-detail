package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type urlResult struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	channel := make(chan urlResult)
	urlList := []string{
		"https://github.com",
		"https://stackoverflow.com",
		"https://google.com",
		"https://youtube.com",
		"http://localhost:8080",
	}

	// create ctx to handle the timeouts
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, url := range urlList {
		go getState(ctx, url, channel)
	}

	for range urlList {
		res := <-channel

		if res.err != nil {
			fmt.Println(res.url, res.latency, res.err)
		} else {
			fmt.Println(res.url, res.latency, res.err)
		}
	}
}

func getState(ctx context.Context, url string, ch chan urlResult) {
	start := time.Now()

	ctxReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(ctxReq); err != nil {
		ch <- urlResult{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- urlResult{url, nil, t}

		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
