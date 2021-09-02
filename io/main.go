package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for _, file_name := range os.Args[1:] {
		var line_count, word_count, char_count int

		// open given file
		file, err := os.Open(file_name)

		if err != nil {
			// show up the errs if it existed.
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// initialize new scanner
		scanner := bufio.NewScanner(file)

		// loop over the lines of the scanner.Scan()
		for scanner.Scan() {
			curr_len := scanner.Text()

			// modify all pre defined vars with the 'curr_len' value
			// curr_len = every line of entire given file.
			line_count++
			char_count += len(curr_len)
			word_count += len(strings.Fields(curr_len))
		}

		fmt.Println("")
		fmt.Println("+------------------------------------------------+")
		fmt.Printf("%-7s | %-7s | %-7s | %s\n", "lines", "words", "chars", "file name")
		fmt.Printf("%-7d | %-7d | %-7d | %s\n", line_count, word_count, char_count, file_name)
		fmt.Println("+------------------------------------------------+")
		file.Close()
	}
}
