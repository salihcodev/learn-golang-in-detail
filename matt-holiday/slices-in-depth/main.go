package main

import "fmt"

func main() {
	a := make([]int, 0, 5) // this the right initialization
	b := make([]int, 5)

	fmt.Println(a) // []
	fmt.Println(b) // [0 0 0 0 0]

	fmt.Println("+-------------------------------------+")

	a2 := append(a, 1)
	b2 := append(b, 1)

	fmt.Println(a2) // [1] coz the capacity with zero
	fmt.Println(b2) // [0 0 0 0 0 1] // at's appending only after the default zero-values.

}
