package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	channel := make(chan result)
	urlList := []string{
		"https://github.com",
		"https://stackoverflow.com",
		"https://google.com",
		"https://asalih.netlify.com",
		"https://youtube.com",
	}

	for _, url := range urlList {
		go get(url, channel)
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

func get(url string, ch chan result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}

		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
