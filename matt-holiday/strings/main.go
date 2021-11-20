package main

import "fmt"

func main() {
	s := "café"

	fmt.Printf("%T %[1]v\n", s)
	fmt.Printf("%T %[1]v\n", len(s))

	// in different langs they are not the same.
	//  []rune || []byte
	fmt.Printf("%T %[1]v\n", []rune(s))
	fmt.Printf("%T %[1]v\n", []byte(s))

}
