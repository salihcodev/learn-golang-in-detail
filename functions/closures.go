package main

import "fmt"

func main() {

	res := closureContainer()
	fmt.Println(res()) // 1
	fmt.Println(res()) // 2
	fmt.Println(res()) // 3

}

// closure ex:
func closureContainer() func() int {
	/*
		those are called env. variables they are 'pointers'

			*x
			*y
	*/
	x, y := 0, 1

	return func() int {
		x, y = y, x+y
		return y
	}
}

func varLifeTime() *int {

	// this variable should be local variable -AKA- not readable outside its function's scop.
	var c int

	// BUT by returning variable's pointer we acceded its lifetime outside the function scope.
	return &c
}
