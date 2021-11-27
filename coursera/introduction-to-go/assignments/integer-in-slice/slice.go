package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {


	var input string
	userHistory :=  make([]int, 3) // create slice with length (3).

	for {
		fmt.Println()
		fmt.Print("Please enter an integer to be added to the slice: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
		}


		if input == "X"  {
			fmt.Println("Exit.")
			break
		}else {

			// convert the input to `int` to can sort it with sort.Ints built in method:
			_input, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
			}

			// append the input to the slice:
			userHistory = append(userHistory, _input)

			// sort user history
			sort.Ints(userHistory)

			fmt.Println()
			fmt.Println("========== Your history upon till now ==========")

			// show the result every round of the loop
			fmt.Println(userHistory)

			fmt.Println("================================================")

		}
	}

}
