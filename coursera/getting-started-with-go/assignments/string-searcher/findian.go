package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {


	/* REQS
	 	starts with 'i', ends with 'n' and contains 'a'
	 	Found! otherwise, Not Found!
		sensitivity
	*/

	fmt.Print("Please enter a string to search in: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		processedInput := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if
		strings.Contains(processedInput, "a") &&
		strings.IndexRune(processedInput, 'i') == 0 &&
		strings.IndexRune(processedInput, 'n') == len(processedInput) - 1 {

			fmt.Println("Found!")

		} else {

			fmt.Println("Not Found!")
		}

		os.Exit(0)
	}
}
