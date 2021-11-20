package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough arguments as it suppose to be.")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println(scanner.Scan()) // this is the first line of the input, so we need to loop

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), old)
		output := strings.Join(input, new)

		fmt.Println(output)
	}
}
