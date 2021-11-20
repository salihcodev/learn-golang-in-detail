package main

import (
	"log"
	"time"
)

func chansFeeder(i int, ch chan<- int) {

	for {
		time.Sleep(time.Duration(i+1) * time.Second)
		ch <- i
	}
}

func main() {
	channels := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range channels {
		go chansFeeder(i, channels[i])
	}

	for {
		select {
		// select re-executing the case's logic every the case has got new value.
		case c0 := <-channels[0]:
			log.Println(c0, "#")
		case c1 := <-channels[1]:
			log.Println(c1, "##")
		}
	}
}
