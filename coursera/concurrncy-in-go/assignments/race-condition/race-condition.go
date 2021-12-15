package main

import (
	"fmt"
	"sync"
)

/*
	Here, while using some sort of data-structure and add data to it via several IIFEs,
	chances are that two functions will try to access the data structure in the exact same time
	AKA 'Race Condition'
*/

func main() {
	var sliceOfThings []string
	wg := &sync.WaitGroup{}

	// I'm using wait group to give every function chance to its job.

	wg.Add(4)

	go func(wg *sync.WaitGroup) {
		sliceOfThings = append(sliceOfThings, "printer")

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		sliceOfThings = append(sliceOfThings, "boat")

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		sliceOfThings = append(sliceOfThings, "goat")

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		sliceOfThings = append(sliceOfThings, "sheet")

		wg.Done()
	}(wg)

	wg.Wait()

	/*
		If you tried to execute this program with --race flag,
		you will figure out how many possibilities for making race condition.
	*/
	fmt.Println("Please try to pass (--race) like [go run --race main.go] flag while executing the program :)\n" +
		"to see possibilities of making a race condition case.")

	// print results
	fmt.Println()
	fmt.Println(sliceOfThings)
}
