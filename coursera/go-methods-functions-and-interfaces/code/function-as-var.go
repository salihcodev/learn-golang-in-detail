package main

import "fmt"


// single arg:
//var funcVar func(int) int

// more than one arg:
var funcVar func(x, y int) int

func testFunc(x, y int) int {
	return x * y
}

func main () {
	funcVar = testFunc

	fmt.Println(funcVar(2, 2))
}
