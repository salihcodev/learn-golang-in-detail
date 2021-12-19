package main

import (
	"fmt"
	"sort"
)

const EndUserMsg = "Please add huge sequence of numbers IDK maybe 40 int, in format: 9,-1,3.... > "

func HandleEndUser(slice []int, err error) []int {
	if err != nil {
		return slice
	}

	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		slice = append(slice, d)
	}

	return HandleEndUser(slice, err)
}

func Swap(data []int, idx int) {
	data[idx], data[idx+1] = data[idx+1], data[idx]
}

func BubbleSort(slice []int, c chan []int) {
	isSwapped := true

	for isSwapped {
		isSwapped = false

		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
				isSwapped = true
			}
		}

	}

	fmt.Println("Slice portion: ", slice)

	// send slice with the channel:
	c <- slice
}

func main() {
	var slice, sortedSlice []int

	//get the user input and store it to be ready for sorting.
	fmt.Println("NOTE: Enter twice")
	fmt.Print(EndUserMsg)
	slice = HandleEndUser(slice, nil)

	// partitioning the slice:
	slicePortion := len(slice) / 4
	s1 := slice[:slicePortion]
	s2 := slice[slicePortion : 2*(slicePortion)]
	s3 := slice[2*(slicePortion) : 3*(slicePortion)]
	s4 := slice[3*(slicePortion):]

	// channels AKA array portions:
	c := make(chan []int)

	go BubbleSort(s1, c)
	go BubbleSort(s2, c)
	go BubbleSort(s3, c)
	go BubbleSort(s4, c)

	for i := 1; i <= 4; i++ {
		// append every current, ready, sorted slice
		sortedSlice = append(sortedSlice, <-c...)
	}

	sort.Ints(slice)
	fmt.Println("The whole sorted slice: ", sortedSlice)
}
