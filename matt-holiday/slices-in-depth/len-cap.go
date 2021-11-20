package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4}

	/*
		disclimare::
		if c:= [i, j, k] then len = j - i and cap = k - i

		b := a[0:1:1]
	*/

	// here we are limiting the length and capacity.
	b := a[0:1:1]
	c := a[0:1] // here it copied the underline cpacity

	fmt.Println("B", b)              // [1]
	fmt.Println("length:B", len(b))  // 1
	fmt.Println("capcity:B", cap(b)) // 1

	fmt.Println("C", c)              // [1]
	fmt.Println("length:C", len(c))  // 1
	fmt.Println("capcity:C", cap(c)) // 4
}
