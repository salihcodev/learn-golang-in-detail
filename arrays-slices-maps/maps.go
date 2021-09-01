package main

import "fmt"

// it's something more like dictionaries in python.
func main() {

	/* create initial map
		- make(map[string]int)
		- map[string]{}

	// create with values
		- map[string]{"me": 5}
	*/

	a := map[string]int{"me": 5}
	b := make(map[string]int)

	c1, ok1 := a["me"] // 5, true | 5 is the value and true coz it's existed.
	c2, ok2 := b["me"] // 0, false | 0 is the initial value and false coz it'snot existed.

	a["me"]++
	c3, ok3 := a["me"] // 6, true

	fmt.Println(c1, ok1)
	fmt.Println(c2, ok2)
	fmt.Println(c3, ok3)

	t := map[string]int{}
	fmt.Println(t == nil) // false, not nil BUT empty.

}
