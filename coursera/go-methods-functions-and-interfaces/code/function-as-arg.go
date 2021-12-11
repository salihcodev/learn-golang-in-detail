package main

import (
	"fmt"
)


//  the passed function to the parent.
func childFunc (val int) {
	fmt.Println(val)
}

// function that takes another function as an argument.
func parentFunc(testFunction func (int), val int) {
	testFunction(val)
}


func main () {
	parentFunc(childFunc, 6)
}