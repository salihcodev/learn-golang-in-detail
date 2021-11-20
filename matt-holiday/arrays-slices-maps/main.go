package main

import "fmt"

var (
	a []int
	b = []int{1, 2, 3}
)

func main() {
	a = append(a, 1) // nil accepts to append
	// fmt.Println(a)   // [1]

	a = b
	fmt.Println(a) // [1,2,3] override 'a' via 'b'

	// 	create a slice via 'make' method:
	c := make([]int, 3)
	fmt.Println(c) // [0 0 0] create given size '3', and has zero-value for every element.

	equality := a[0] == b[0]
	fmt.Println(equality) // true
}
