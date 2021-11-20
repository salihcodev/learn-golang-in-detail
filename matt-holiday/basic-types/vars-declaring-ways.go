package main

import "fmt"

/*
	declaring variables outside the main function:
*/
// #1
var a int = 7

// #2
var (
	c string = "hello" // get rid of `string` no problems
	d        = "world"
	f        = 1.5 // here no need for identifying the type `which is float64`
)

const (
	g = "string"
)

func main() {
	/*
		declaring variables inside the main function:
	*/

	b := 5

	fmt.Println(a, b)
	fmt.Println(c, d)

	// get some info about the variable via terminal
	fmt.Printf("b: %T %[1]v\n", b)

}
