package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	words := map[string]int{}

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	fmt.Println("+--------------------------------------------------------+")
	for k, v := range words {

		fmt.Printf("|%-15s |%5d\n", k, v)
		fmt.Println("|--------------------------------------------------------|")
	}
	fmt.Println("+--------------------------------------------------------+")

}
