package main

import "fmt"

const EndUserMsg = "Please add sequence of numbers to sort like 9,-1,3: "


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

func main() {

	var slice []int

	// get the user input and store it to be ready for sorting.
	fmt.Println("NOTE: Enter twice")
	fmt.Print(EndUserMsg)
	slice = HandleEndUser(slice, nil)

	// sort the given slice
	BubbleSort(slice)

	fmt.Println("Output In Sorted Way: ", slice)
}