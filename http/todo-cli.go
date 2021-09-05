package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseUrl = "https://jsonplaceholder.typicode.com"

type todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	/*
		fetch todos depending on user input `from CLI`
			- get user input:  #1
			- concatinating it to baseUrl:  #2
	*/

	// #1
	todoID := os.Args[1]

	// #2
	res, err := http.Get(baseUrl + "/todos/" + todoID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer res.Body.Close()
	var todoI todo

	if res.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		err = json.Unmarshal(body, &todoI)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

	}

	// printing the result:
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println(todoI)
	fmt.Println("")
	fmt.Printf("%#v\n", todoI)
	fmt.Println("#-----------------------------------------------------#")
}
