package main

import (
	"fmt"
	"strconv"
)

func HandleEndUser(slice *[]int) {
	var input string

	// get the numbers sequence
	fmt.Print("Please add sequence of numbers to sort like 361: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}

	// loop over the string input
	for _, val := range input {
		// convert string to int
		elm, err := strconv.Atoi(string(val))
		if err != nil {
			fmt.Println(err)
		}

		// append converted element, to the slice
		*slice = append(*slice, elm)
	}
}

func Swap(data []int, idx int) {
	temp := data[idx]
	data[idx] =  data[idx+1]
	data[idx+1] =  temp

	// another solution:
	//data[idx], data[idx+1] = data[idx+1], data[idx]
}

func BubbleSort(slice []int) {
	isSwapped:= true

	for isSwapped {
		/*
			make swap flag to check every time, is the swap happened or not
			coz if it not, then there is no need for complete sorting the given data.
		*/
		isSwapped = false

		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
				isSwapped = true
			}
		}
	}
}



func main() {
	var sliceToSort []int

	// get the user input and store it to be ready for sorting.
	HandleEndUser(&sliceToSort)

	// sort the given slice
	BubbleSort(sliceToSort)

	fmt.Println("Output In Sorted Way: ", sliceToSort)
}
