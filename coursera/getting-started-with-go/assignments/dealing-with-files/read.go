package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type line struct {
	fname string
	lname string
}

func main() {
	var filename string

	fmt.Print("Please enter the file name: ")
	_, err := fmt.Scan(&filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " ")

		var sl []line

		l := line{t[0], t[1]}
		sl = append(sl, l)


		for _, val :=  range sl {
			fmt.Printf("Fisrt Name: %s  | Last Name: %s\n", val.fname, val.lname)
		}
	}
}
