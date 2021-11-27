package main

import ("fmt"
)

func main() {
	var userInput float32


	// Get the value of the user input, Set it to userInput variable:
	fmt.Scanln(&userInput)

	// Print the original, And convert it to integer.
	fmt.Printf("Output: %g | Converted output to integer: %d\n", userInput, int(userInput))



}
