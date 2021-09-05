package main

import (
	"fmt"
	"log"
	"net/http"
)

// first web-server
func main() {

	http.HandleFunc("/", basicHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from %s\n", r.URL.Path[1:])
}
