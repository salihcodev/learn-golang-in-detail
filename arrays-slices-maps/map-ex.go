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

	for k, v := range words {

		fmt.Println("Word [", k, "]: appeared ---------------------------- (", v, ") times")
	}

}
