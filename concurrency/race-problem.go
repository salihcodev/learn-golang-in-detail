package main

import (
	"fmt"
	"log"
	"net/http"
)

type chanVal chan int

// 5) fire the server up, return the channel value back.
func (ch chanVal) server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Result: %d`, <-ch)
}

// 3) increase the channel value:
func calc(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	// 1) creating the channel:
	var ch chanVal = make(chan int)

	// 2) pass the channel to calc function:
	go calc(ch)

	// 4) make sure that the channel has the same type to run thr method `server` having the channel value:
	http.HandleFunc("/", ch.server)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
