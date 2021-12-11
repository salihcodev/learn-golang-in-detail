package main

import "fmt"


// variadic function
func variadicFunction(nums ...int) int {
	max := -1
	for _, elm := range nums {

		if elm > max {
			max = elm
		}
	}

	return max
}

func main () {

	fmt.Println("Max num is: ", variadicFunction(2,0,99,1,5,8))


	// pass a punch of integers as a slice, slice... means 'pre-packaged'
	slice := []int{2,0,79,1,5,8}
	fmt.Println("Max num is: ", variadicFunction(slice...))
}
