package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// first web-server
func main() {

	http.HandleFunc("/", basicHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(10) * time.Second) // this line just for testing ctx handler.
	fmt.Fprintf(w, "Hello world from %s\n", r.URL.Path[1:])
}
